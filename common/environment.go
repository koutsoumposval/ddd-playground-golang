package environment

import (
	"os"
)

func getEnvVar(name, def string) string {
	env := os.Getenv(name)

	if len(env) != 0 {
		return env
	}

	return def
}
