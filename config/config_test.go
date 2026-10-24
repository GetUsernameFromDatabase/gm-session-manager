package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func lookupMap(m map[string]string) lookupFunc {
	return func(key string) (string, bool) {
		v, ok := m[key]
		return v, ok
	}
}

func TestFromLookup(t *testing.T) {
	valid := map[string]string{
		EnvAdminUsername: "gm",
		EnvAdminPassword: "secret",
	}
	with := func(overrides map[string]string) map[string]string {
		m := map[string]string{}
		for k, v := range valid {
			m[k] = v
		}
		for k, v := range overrides {
			m[k] = v
		}
		return m
	}

	tests := []struct {
		name    string
		env     map[string]string
		want    Config
		wantErr []string
	}{
		{
			name: "defaults applied",
			env:  valid,
			want: Config{AdminUsername: "gm", AdminPassword: "secret", Port: DefaultPort, DataDir: DefaultDataDir},
		},
		{
			name: "explicit values",
			env:  with(map[string]string{EnvPort: "9000", EnvDataDir: "/srv/gm"}),
			want: Config{AdminUsername: "gm", AdminPassword: "secret", Port: 9000, DataDir: "/srv/gm"},
		},
		{name: "missing username", env: with(map[string]string{EnvAdminUsername: ""}), wantErr: []string{EnvAdminUsername}},
		{name: "missing password", env: with(map[string]string{EnvAdminPassword: ""}), wantErr: []string{EnvAdminPassword}},
		{name: "whitespace password", env: with(map[string]string{EnvAdminPassword: "   "}), wantErr: []string{EnvAdminPassword}},
		{name: "both missing", env: map[string]string{}, wantErr: []string{EnvAdminUsername, EnvAdminPassword}},
		{name: "port not a number", env: with(map[string]string{EnvPort: "http"}), wantErr: []string{EnvPort}},
		{name: "port out of range", env: with(map[string]string{EnvPort: "70000"}), wantErr: []string{EnvPort}},
		{name: "port zero", env: with(map[string]string{EnvPort: "0"}), wantErr: []string{EnvPort}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fromLookup(lookupMap(tt.env))
			if len(tt.wantErr) > 0 {
				if err == nil {
					t.Fatalf("expected error mentioning %v, got config %+v", tt.wantErr, got)
				}
				for _, want := range tt.wantErr {
					if !strings.Contains(err.Error(), want) {
						t.Errorf("error %q does not mention %s", err, want)
					}
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLoadEnvFileAndPrecedence(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".env")
	content := "ADMIN_USERNAME=filedm\nADMIN_PASSWORD=filepass\nPORT=8123\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(EnvAdminUsername, "envdm")

	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.AdminUsername != "envdm" {
		t.Errorf("process env should win: got %q", cfg.AdminUsername)
	}
	if cfg.AdminPassword != "filepass" || cfg.Port != 8123 {
		t.Errorf("file values not loaded: %+v", cfg)
	}
}

func TestLoadMissingFileIsNotAnError(t *testing.T) {
	t.Setenv(EnvAdminUsername, "gm")
	t.Setenv(EnvAdminPassword, "pw")
	if _, err := Load(filepath.Join(t.TempDir(), "absent.env")); err != nil {
		t.Fatalf("missing .env should fall back to environment: %v", err)
	}
}
