// Package config loads runtime settings from the environment and an optional
// .env file, and refuses to continue when required values are missing.
package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// Environment variable names.
const (
	EnvAdminUsername = "ADMIN_USERNAME"
	EnvAdminPassword = "ADMIN_PASSWORD"
	EnvPort          = "PORT"
	EnvDataDir       = "DATA_DIR"
)

// Defaults for optional settings.
const (
	DefaultPort    = 8080
	DefaultDataDir = "data"
	// DefaultPlayerDataDir differs so a server and a player app run from one folder never share a database.
	DefaultPlayerDataDir = "data-player"
	// DefaultPlayerPort differs from DefaultPort so the player app can run beside a server.
	DefaultPlayerPort = 8042
	// maxPort matches the max= on Config.Port.
	maxPort = 65535
)

// Config is the validated runtime configuration. The env tag names the
// variable each field comes from and is used in validation messages.
type Config struct {
	// Player is the player app (--player): no admin credentials needed.
	Player        bool
	AdminUsername string `env:"ADMIN_USERNAME" validate:"required_unless=Player true"`
	AdminPassword string `env:"ADMIN_PASSWORD" validate:"required_unless=Player true"`
	Port          int    `env:"PORT"           validate:"min=1,max=65535"`
	DataDir       string `env:"DATA_DIR"       validate:"required"`
}

// lookupFunc matches os.LookupEnv.
type lookupFunc func(key string) (string, bool)

// Load reads envFile (if it exists) and the process environment, with real
// environment variables taking precedence over the file.
func Load(envFile string) (Config, error) {
	fileVars, err := godotenv.Read(envFile)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return Config{}, fmt.Errorf("reading %s: %w", envFile, err)
	}
	return fromLookup(func(key string) (string, bool) {
		if v, ok := os.LookupEnv(key); ok {
			return v, true
		}
		v, ok := fileVars[key]
		return v, ok
	})
}

// Player builds the player app's config from flags only: it reads no .env or environment.
func Player(port int, dataDir string) (Config, error) {
	// Checked here, not by the validator: its messages name .env variables.
	if port < 1 || port > maxPort {
		return Config{}, fmt.Errorf("config: --port must be 1-%d", maxPort)
	}
	if strings.TrimSpace(dataDir) == "" {
		return Config{}, errors.New("config: --data must not be empty")
	}
	return Config{Player: true, Port: port, DataDir: dataDir}, nil
}

func fromLookup(lookup lookupFunc) (Config, error) {
	get := func(key, fallback string) string {
		if v, ok := lookup(key); ok && strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
		return fallback
	}

	cfg := Config{
		AdminUsername: get(EnvAdminUsername, ""),
		AdminPassword: get(EnvAdminPassword, ""),
		DataDir:       get(EnvDataDir, DefaultDataDir),
	}

	port, err := strconv.Atoi(get(EnvPort, strconv.Itoa(DefaultPort)))
	if err != nil {
		return Config{}, fmt.Errorf("config: %s must be a number", EnvPort)
	}
	cfg.Port = port

	if err := validate.Struct(cfg); err != nil {
		return Config{}, describe(err)
	}
	return cfg, nil
}

var validate = newValidator()

// newValidator reports field errors by their env tag so messages name the
// variable the user has to set.
func newValidator() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterTagNameFunc(func(f reflect.StructField) string {
		return f.Tag.Get("env")
	})
	return v
}

func describe(err error) error {
	var fieldErrs validator.ValidationErrors
	if !errors.As(err, &fieldErrs) {
		return fmt.Errorf("config: %w", err)
	}
	msgs := make([]string, 0, len(fieldErrs))
	for _, fe := range fieldErrs {
		switch fe.Tag() {
		case "required", "required_unless":
			msgs = append(msgs, fe.Field()+" is required (set it in .env, see .env.example)")
		default:
			msgs = append(msgs, fmt.Sprintf("%s is invalid (%s=%s)", fe.Field(), fe.Tag(), fe.Param()))
		}
	}
	return errors.New("config: " + strings.Join(msgs, "; "))
}
