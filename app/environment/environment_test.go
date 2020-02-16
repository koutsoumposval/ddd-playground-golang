package environment

import "testing"

func TestGetEnvVarReturnDefaultValue(t *testing.T) {
	expected := "Default value"

	result := GetEnvVar("NON_EXISTENT_ENV_VAR", expected)

	if result != expected {
		t.Errorf("GetEnvVar() test failed. Result %q, expected %q", result, expected)
	}
}
