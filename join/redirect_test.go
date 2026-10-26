package join

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

// serveSplit serves app over TLS and the redirect over plain HTTP on one
// port, and returns its address and a client trusting its certificate.
func serveSplit(t *testing.T) (string, *http.Client) {
	t.Helper()
	certPEM, keyPEM, err := generate(nil, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		t.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(certPEM)

	base, err := (&net.ListenConfig{}).Listen(t.Context(), "tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	tlsL, plainL := SplitTLS(base)
	app := &http.Server{
		Handler:           http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _, _ = io.WriteString(w, "app") }),
		ReadHeaderTimeout: time.Second,
		TLSConfig:         &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12},
	}
	redirect := &http.Server{Handler: RedirectHTTPS, ReadHeaderTimeout: time.Second}
	go func() { _ = app.ServeTLS(tlsL, "", "") }()
	go func() { _ = redirect.Serve(plainL) }()
	t.Cleanup(func() {
		_ = app.Close()
		_ = redirect.Close()
		_ = base.Close()
	})

	client := &http.Client{
		Transport:     &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool, MinVersion: tls.VersionTLS12}},
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       5 * time.Second,
	}
	return base.Addr().String(), client
}

func TestSplitTLS(t *testing.T) {
	addr, client := serveSplit(t)
	tests := []struct {
		name         string
		method, url  string
		wantStatus   int
		wantLocation string
		wantBody     string
	}{
		{"plain GET redirects", http.MethodGet, "http://" + addr + "/notes?x=1", http.StatusPermanentRedirect, "https://" + addr + "/notes?x=1", ""},
		{"plain POST redirects", http.MethodPost, "http://" + addr + "/login", http.StatusPermanentRedirect, "https://" + addr + "/login", ""},
		{"TLS reaches the app", http.MethodGet, "https://" + addr + "/", http.StatusOK, "", "app"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = resp.Body.Close() }()
			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode != tt.wantStatus || resp.Header.Get("Location") != tt.wantLocation {
				t.Errorf("got %d %q, want %d %q", resp.StatusCode, resp.Header.Get("Location"), tt.wantStatus, tt.wantLocation)
			}
			if tt.wantBody != "" && string(body) != tt.wantBody {
				t.Errorf("body %q, want %q", body, tt.wantBody)
			}
		})
	}
}

// A connection that never speaks is closed instead of held open.
func TestSplitTLSSilentConn(t *testing.T) {
	old := peekTimeout
	peekTimeout = 50 * time.Millisecond
	t.Cleanup(func() { peekTimeout = old })
	addr, _ := serveSplit(t)

	c, err := (&net.Dialer{}).DialContext(t.Context(), "tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }()
	_ = c.SetReadDeadline(time.Now().Add(2 * time.Second))
	if _, err := c.Read(make([]byte, 1)); !errors.Is(err, io.EOF) {
		t.Errorf("read on a silent connection: %v, want EOF (closed by the server)", err)
	}
}
