package tools

import (
	"testing"
)

func TestBuildResponse(t *testing.T) {
	want := map[string]interface{}{
		"message": "test",
	}
	got, err := BuildResponse(200, want)
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

func TestStringify(t *testing.T) {
	want := map[string]interface{}{
		"message": "test",
	}
	got := Stringify(want)
	if got != `{"message":"test"}` {
		t.Errorf("Response = %q, want %q", got, `{"message":"test"}`)
	}
}
