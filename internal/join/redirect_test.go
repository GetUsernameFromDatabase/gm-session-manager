package join

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	routes "gm-session-manager/internal/route" // join has its own route func
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
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, _ *http.Request) { _, _ = io.WriteString(w, "app") },
		),
		ReadHeaderTimeout: time.Second,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
		},
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
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool, MinVersion: tls.VersionTLS12},
		},
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
		{
			"plain GET redirects",
			http.MethodGet,
			"http://" + addr + routes.Notes + "?x=1",
			http.StatusPermanentRedirect,
			"https://" + addr + routes.Notes + "?x=1",
			"",
		},
		{
			"plain POST redirects",
			http.MethodPost,
			"http://" + addr + routes.AdminLogin,
			http.StatusPermanentRedirect,
			"https://" + addr + routes.AdminLogin,
			"",
		},
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
				t.Errorf(
					"got %d %q, want %d %q",
					resp.StatusCode,
					resp.Header.Get("Location"),
					tt.wantStatus,
					tt.wantLocation,
				)
			}
			if tt.wantBody != "" && string(body) != tt.wantBody {
				t.Errorf("body %q, want %q", body, tt.wantBody)
			}
		})
	}
}

// The redirect only echoes hosts this machine can be reached by.
func TestRedirectHTTPSHost(t *testing.T) {
	own, _ := os.Hostname()
	tests := []struct {
		host string
		want int
	}{
		{"evil.example", http.StatusBadRequest},
		{"evil.example:8443", http.StatusBadRequest},
		{"192.168.1.20:8443", http.StatusPermanentRedirect},
		{"[::1]:8443", http.StatusPermanentRedirect},
		{"localhost:8443", http.StatusPermanentRedirect},
		{strings.ToUpper(own) + ":8443", http.StatusPermanentRedirect},
		{own + ".local:8443", http.StatusPermanentRedirect},
		{own + ".evil.example", http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.host, func(t *testing.T) {
			req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, routes.Notes, nil)
			req.Host = tt.host
			rec := httptest.NewRecorder()
			RedirectHTTPS.ServeHTTP(rec, req)
			if rec.Code != tt.want {
				t.Errorf("got %d, want %d", rec.Code, tt.want)
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
