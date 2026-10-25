package hub

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

const (
	// helloTimeout: a client that never says hello is synced from now on,
	// without catch-up.
	helloTimeout = 5 * time.Second
	pingInterval = 30 * time.Second
	writeTimeout = 10 * time.Second
	// ReadLimit caps one incoming message, decompressed.
	ReadLimit = 64 << 10
)

// clientMessage is what browsers send over the socket.
type clientMessage struct {
	Type string `json:"type"`
	// hello
	Epoch string `json:"epoch"`
	Seq   uint64 `json:"seq"`
	// hello: the live views the page shows
	Watch []string `json:"watch"`
	// sub / unsub
	Blocks []int64 `json:"blocks"`
}

// Serve upgrades the request to a websocket for c and runs it until the
// connection ends. The caller has already authenticated the request; the
// default origin check rejects cross-site pages.
func (h *Hub) Serve(w http.ResponseWriter, r *http.Request, c *Client) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		slog.Debug("websocket accept", "err", err)
		return
	}
	defer func() { _ = conn.CloseNow() }()
	conn.SetReadLimit(ReadLimit)

	h.register(c)
	defer h.unregister(c)

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	go func() {
		select {
		case <-c.done:
		case <-ctx.Done():
		}
		cancel()
	}()
	go h.readLoop(ctx, conn, c)
	go func() {
		select {
		case <-time.After(helloTimeout):
			h.mu.Lock()
			synced := c.synced
			h.mu.Unlock()
			if !synced {
				_, seq := h.Position()
				h.sync(c, h.epoch, seq)
			}
		case <-ctx.Done():
		}
	}()

	err = h.writeLoop(ctx, conn, c)
	if err != nil && !errors.Is(err, context.Canceled) && websocket.CloseStatus(err) == -1 {
		slog.Debug("websocket closed", "err", err)
	}
}

func (h *Hub) readLoop(ctx context.Context, conn *websocket.Conn, c *Client) {
	defer c.stop()
	for {
		_, data, err := conn.Read(ctx)
		if err != nil {
			return
		}
		var m clientMessage
		if json.Unmarshal(data, &m) != nil {
			continue
		}
		switch m.Type {
		case "hello":
			c.setWatch(m.Watch)
			h.sync(c, m.Epoch, m.Seq)
		case "sub":
			for _, id := range m.Blocks {
				if h.handler.CanSubscribe(ctx, c, id) {
					c.subscribe(id, true)
				}
			}
		case "unsub":
			for _, id := range m.Blocks {
				c.subscribe(id, false)
			}
		}
	}
}

func (h *Hub) writeLoop(ctx context.Context, conn *websocket.Conn, c *Client) error {
	ping := time.NewTicker(pingInterval)
	defer ping.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ping.C:
			pctx, cancel := context.WithTimeout(ctx, writeTimeout)
			err := conn.Ping(pctx)
			cancel()
			if err != nil {
				return err
			}
		case m := <-c.out:
			payload, err := h.payload(ctx, c, m)
			if err != nil {
				slog.Error("render websocket update", "err", err)
				continue
			}
			if len(payload) > 0 {
				wctx, cancel := context.WithTimeout(ctx, writeTimeout)
				err = conn.Write(wctx, websocket.MessageText, payload)
				cancel()
				if err != nil {
					return err
				}
			}
			if m.close {
				return conn.Close(websocket.StatusNormalClosure, "session ended")
			}
		}
	}
}

// payload encodes one queued message.
func (h *Hub) payload(ctx context.Context, c *Client, m message) ([]byte, error) {
	switch {
	case m.raw != nil:
		return m.raw, nil
	case m.control != "":
		return json.Marshal(map[string]string{"type": m.control})
	}
	return h.patch(ctx, c, m)
}
