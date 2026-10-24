package join

import (
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var (
	lanA = net.ParseIP("192.168.1.20")
	lanB = net.ParseIP("10.0.0.5")
	t0   = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
)

func TestLoadOrCreateCert(t *testing.T) {
	tests := []struct {
		name        string
		seed        func(t *testing.T, dir string) // prepares dir before the call
		ips         []net.IP
		now         time.Time
		wantCreated bool
	}{
		{"missing files", nil, []net.IP{lanA}, t0, true},
		{"valid and covering", seedCert([]net.IP{lanA}), []net.IP{lanA}, t0.Add(24 * time.Hour), false},
		{"covers a superset", seedCert([]net.IP{lanA, lanB}), []net.IP{lanA}, t0, false},
		{"new LAN address", seedCert([]net.IP{lanA}), []net.IP{lanA, lanB}, t0, true},
		{"close to expiry", seedCert([]net.IP{lanA}), []net.IP{lanA}, t0.Add(certValidity - renewBefore + time.Hour), true},
		{"expired", seedCert([]net.IP{lanA}), []net.IP{lanA}, t0.Add(certValidity + time.Hour), true},
		{"corrupt files", seedGarbage, []net.IP{lanA}, t0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			if tt.seed != nil {
				tt.seed(t, dir)
			}
			cert, created, err := LoadOrCreateCert(dir, tt.ips, tt.now)
			if err != nil {
				t.Fatal(err)
			}
			if created != tt.wantCreated {
				t.Errorf("created = %v, want %v", created, tt.wantCreated)
			}
			if !usable(cert.Leaf, tt.ips, tt.now) {
				t.Error("returned certificate does not cover the requested addresses at now")
			}
			// Whatever was returned must be what is on disk, so the next start reuses it.
			again, created, err := LoadOrCreateCert(dir, tt.ips, tt.now)
			if err != nil {
				t.Fatal(err)
			}
			if created || Fingerprint(again) != Fingerprint(cert) {
				t.Error("second call did not reuse the stored certificate")
			}
		})
	}
}

func TestCertNames(t *testing.T) {
	cert, _, err := LoadOrCreateCert(t.TempDir(), []net.IP{lanA}, t0)
	if err != nil {
		t.Fatal(err)
	}
	for _, host := range []string{"localhost", "127.0.0.1", "::1", lanA.String()} {
		if err := cert.Leaf.VerifyHostname(host); err != nil {
			t.Errorf("%s: %v", host, err)
		}
	}
	if err := cert.Leaf.VerifyHostname(lanB.String()); err == nil {
		t.Errorf("%s should not be covered", lanB)
	}
}

func TestKeyFileIsPrivate(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("Unix permission bits are not meaningful on Windows")
	}
	dir := t.TempDir()
	if _, _, err := LoadOrCreateCert(dir, nil, t0); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(filepath.Join(dir, KeyFile))
	if err != nil {
		t.Fatal(err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Errorf("key file mode = %o, want 600", perm)
	}
}

func seedCert(ips []net.IP) func(t *testing.T, dir string) {
	return func(t *testing.T, dir string) {
		t.Helper()
		if _, _, err := LoadOrCreateCert(dir, ips, t0); err != nil {
			t.Fatal(err)
		}
	}
}

func seedGarbage(t *testing.T, dir string) {
	t.Helper()
	for _, f := range []string{CertFile, KeyFile} {
		if err := os.WriteFile(filepath.Join(dir, f), []byte("not pem"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
}
