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
)

// Config is the validated runtime configuration. The env tag names the
// variable each field comes from and is used in validation messages.
type Config struct {
	AdminUsername string `env:"ADMIN_USERNAME" validate:"required"`
	AdminPassword string `env:"ADMIN_PASSWORD" validate:"required"`
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
		case "required":
			msgs = append(msgs, fe.Field()+" is required (set it in .env, see .env.example)")
		default:
			msgs = append(msgs, fmt.Sprintf("%s is invalid (%s=%s)", fe.Field(), fe.Tag(), fe.Param()))
		}
	}
	return errors.New("config: " + strings.Join(msgs, "; "))
}
