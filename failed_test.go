package tools

import (
	"testing"
)

func TestFailed(t *testing.T) {
	want := "test"

	got, err := Failed(want);
	if err != nil {
		t.Errorf("Error = %q, want nil", err.Error())
	}

	if got.StatusCode != 400 {
		t.Errorf("StatusCode = %q, want %q", got.StatusCode, 400)
	}

	if got.Body != Stringify(map[string]interface{}{"message": want}) {
		t.Errorf("Body = %q, want %q", got.Body, Stringify(map[string]interface{}{"message": want}))
	}

}
