package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "ApiKey testtoken")

	key, err := GetAPIKey(req.Header)
	if err != nil {
		t.Errorf("Fail with error %v", err)
	}
	want := "testtoken"
	if key != want {
		t.Fail()
	}
}
