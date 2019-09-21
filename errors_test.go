package golang_serverless_response

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	want := "test"
	got, err := Errors(errors.New(want));
	if err != nil {
		t.Errorf("Error = %q, want nil", err.Error())
	}
	if got.StatusCode != 500 {
		t.Errorf("StatusCode = %q, want %q", got.StatusCode, 500)
	}

	if got.Body != Stringify(map[string]interface{}{"message": want}) {
		t.Errorf("Body = %q, want %q", got.Body, Stringify(map[string]interface{}{"message": want}))
	}
}
