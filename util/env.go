package util

import (
	"os"
	"strings"
)

// Env returns a OS env var or default value
func Env(key string, defvalue string) string {
	env := os.Getenv(key)

	if env == "" {
		env = defvalue
	}

	if !strings.HasPrefix(env, ":") {
		env = ":" + env
	}

	return env
}
