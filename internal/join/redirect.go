package join

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"
)

// tlsHandshake is the first byte of every TLS connection (a handshake record).
const tlsHandshake = 0x16

// peekTimeout bounds how long a new connection may stay silent.
var peekTimeout = 10 * time.Second

// acceptRetry is the pause after a failed Accept before the next try.
const acceptRetry = 50 * time.Millisecond

// SplitTLS serves one port for both: connections that start a TLS handshake
// come out of tlsL, anything else (plain HTTP) out of plainL. Closing l stops
// both.
func SplitTLS(l net.Listener) (tlsL, plainL net.Listener) {
	t, p := newConnListener(l.Addr()), newConnListener(l.Addr())
	go func() {
		defer func() { _, _ = t.Close(), p.Close() }()
		for {
			c, err := l.Accept()
			if errors.Is(err, net.ErrClosed) {
				return
			}
			if err != nil {
				// Passing errors (too many open files) must not end the server,
				// as http.Server retries them too.
				time.Sleep(acceptRetry)
				continue
			}
			go route(c, t, p)
		}
	}()
	return t, p
}

// route peeks at c's first byte and hands c to the matching listener.
func route(c net.Conn, t, p *connListener) {
	r := bufio.NewReader(c)
	_ = c.SetReadDeadline(time.Now().Add(peekTimeout))
	b, err := r.Peek(1)
	_ = c.SetReadDeadline(time.Time{})
	if err != nil {
		_ = c.Close()
		return
	}
	to := p
	if b[0] == tlsHandshake {
		to = t
	}
	to.hand(peekedConn{Conn: c, r: r})
}

// RedirectHTTPS sends a plain HTTP request to the same address over HTTPS.
// 308 keeps the method, so a form posted over HTTP is posted again.
var RedirectHTTPS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusPermanentRedirect)
})

// peekedConn reads the bytes route peeked before the rest of the connection.
type peekedConn struct {
	net.Conn
	r *bufio.Reader
}

func (c peekedConn) Read(b []byte) (int, error) { return c.r.Read(b) }

// connListener is a net.Listener fed connections by SplitTLS.
type connListener struct {
	addr  net.Addr
	conns chan net.Conn
	done  chan struct{}
	once  sync.Once
}

func newConnListener(addr net.Addr) *connListener {
	return &connListener{addr: addr, conns: make(chan net.Conn), done: make(chan struct{})}
}

func (l *connListener) hand(c net.Conn) {
	select {
	case l.conns <- c:
	case <-l.done:
		_ = c.Close()
	}
}

func (l *connListener) Accept() (net.Conn, error) {
	select {
	case c := <-l.conns:
		return c, nil
	case <-l.done:
		return nil, net.ErrClosed
	}
}

func (l *connListener) Close() error {
	l.once.Do(func() { close(l.done) })
	return nil
}

func (l *connListener) Addr() net.Addr { return l.addr }
