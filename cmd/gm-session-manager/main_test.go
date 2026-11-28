package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gm-session-manager/internal/config"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantPlayer bool
		wantPort   int
		wantData   string
		wantErr    bool
	}{
		{
			name:     "no flags is the GM server",
			args:     nil,
			wantPort: config.DefaultPlayerPort,
			wantData: config.DefaultPlayerDataDir,
		},
		{
			name:       "player defaults",
			args:       []string{"--player"},
			wantPlayer: true,
			wantPort:   8042,
			wantData:   "data-player",
		},
		{
			name:       "player overrides",
			args:       []string{"--player", "--port", "9000", "--data", "x"},
			wantPlayer: true,
			wantPort:   9000,
			wantData:   "x",
		},
		{name: "port alone refused", args: []string{"--port", "9000"}, wantErr: true},
		{name: "data alone refused", args: []string{"--data", "x"}, wantErr: true},
		{name: "no-browser alone refused", args: []string{"--no-browser"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player, _, port, data, err := parseFlags(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if player != tt.wantPlayer || port != tt.wantPort || data != tt.wantData {
				t.Errorf("got %v %d %q", player, port, data)
			}
		})
	}
	// A separate loop: a new wantNoBrowser column would touch every row above.
	for _, tt := range []struct {
		args []string
		want bool
	}{
		{[]string{"--player"}, false},
		{[]string{"--player", "--no-browser"}, true},
	} {
		_, noBrowser, _, _, err := parseFlags(tt.args)
		if err != nil || noBrowser != tt.want {
			t.Errorf("%v: noBrowser %v, err %v", tt.args, noBrowser, err)
		}
	}
}

func TestLoopbackHostOnly(t *testing.T) {
	h := loopbackHostOnly(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	tests := []struct {
		host string
		want int
	}{
		{"127.0.0.1:8080", http.StatusOK},
		{"localhost:8080", http.StatusOK},
		{"127.0.0.1", http.StatusOK},
		{"evil.example:8080", http.StatusMisdirectedRequest},
		{"evil.example", http.StatusMisdirectedRequest},
		{"127.0.0.1.evil.example:8080", http.StatusMisdirectedRequest},
		{"", http.StatusMisdirectedRequest},
	}
	for _, tt := range tests {
		r := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
		r.Host = tt.host
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		if w.Code != tt.want {
			t.Errorf("host %q: status %d, want %d", tt.host, w.Code, tt.want)
		}
	}
}
