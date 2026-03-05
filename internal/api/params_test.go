package api

import (
	"net/url"
	"testing"
)

func TestBuildParamsUrl_ContainsAllKeys(t *testing.T) {
	params := map[string]string{
		"action": "query",
		"format": "json",
		"titles": "Joker",
	}

	result := BuildParamsUrl(params)

	parsed, err := url.ParseQuery(result[1:]) // strip leading '?'
	if err != nil {
		t.Fatalf("BuildParamsUrl returned unparseable query string: %v", err)
	}

	for key, expected := range params {
		if got := parsed.Get(key); got != expected {
			t.Errorf("param %q: expected %q, got %q", key, expected, got)
		}
	}
}

func TestBuildParamsUrl_StartsWithQuestionMark(t *testing.T) {
	result := BuildParamsUrl(map[string]string{"a": "b"})
	if result[0] != '?' {
		t.Errorf("expected result to start with '?', got %q", string(result[0]))
	}
}

func TestBuildParamsUrl_EmptyParams(t *testing.T) {
	result := BuildParamsUrl(map[string]string{})
	if result != "?" {
		t.Errorf("expected '?' for empty params, got %q", result)
	}
}

func TestBuildParamsUrl_SpecialCharactersAreEncoded(t *testing.T) {
	params := map[string]string{
		"titles": "Tarot Card",
	}
	result := BuildParamsUrl(params)
	// spaces must be percent-encoded or '+' encoded — raw space is invalid
	if result == "?titles=Tarot Card" {
		t.Error("special characters were not encoded in the URL")
	}
}