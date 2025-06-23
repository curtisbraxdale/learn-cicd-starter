package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expectedOutput := "8jv083qj0q83j4f345t"
	testRequest, _ := http.NewRequest("GET", ".", nil)
	testRequest.Header.Set("Authorization", "ApiKey 8jv083qj0q83j4f345t")
	actualOutput, err := GetAPIKey(testRequest.Header)
	if err != nil {
		t.Fatalf("Error getting API key.")
	}

	if actualOutput != expectedOutput {
		t.Fatalf("GetAPIKey = %q; want %q", actualOutput, expectedOutput)
	}
}

func TestGetAPIKeyEmpty(t *testing.T) {
	testRequest, _ := http.NewRequest("GET", ".", nil)
	_, err := GetAPIKey(testRequest.Header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("Failed empty authorization header test.")
	}
}

func TestGetAPIKeyMalformed(t *testing.T) {
	testRequest, _ := http.NewRequest("GET", ".", nil)
	testRequest.Header.Set("Authorization", "Malformed Auth String")
	_, err := GetAPIKey(testRequest.Header)
	if err != ErrMalformedAuthHeader {
		t.Fatalf("Failed malformed authorization header test.")
	}
}
