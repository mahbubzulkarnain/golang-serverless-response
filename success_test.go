package golang_serverless_response

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	want := map[string]interface{}{"message": "test"}

	got, err := Success(want);
	if err != nil {
		t.Errorf("Error = %q, want nil", err.Error())
	}

	if got.StatusCode != 200 {
		t.Errorf("StatusCode = %q, want %q", got.StatusCode, 200)
	}

	if got.Body != Stringify(want) {
		t.Errorf("Body = %q, want %q", got.Body, Stringify(want))
	}
}
