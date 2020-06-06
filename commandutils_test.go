package go_utils

import "testing"

func TestCommandOutput(t *testing.T) {
	output, err := ExecuteCommandAndGetResults("echo \"hello\"")
	if err != nil {
		t.Error(err)
	}

	if output != "hello" {
		t.Errorf("\"%s\" should be \"%s\"", output, "hello")
	}
}
