package environment

import (
	"os"
)

// GetEnvVar retrieves ENV_VAR value or default
func GetEnvVar(name string, def string) string {
	env := os.Getenv(name)

	if len(env) != 0 {
		return env
	}

	return def
}
