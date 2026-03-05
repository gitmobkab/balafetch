package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestGetRequest_Success_ReturnsBody(t *testing.T) {
	expected := `{"query": "result"}`
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(expected))
	}))
	defer server.Close()

	body, err := GetRequest(server.URL, 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(body) != expected {
		t.Errorf("expected body %q, got %q", expected, string(body))
	}
}

func TestGetRequest_Non2xxStatus_ReturnsError(t *testing.T) {
	for _, status := range []int{400, 404, 500, 503} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
			}))
			defer server.Close()

			_, err := GetRequest(server.URL, 5)
			if err == nil {
				t.Errorf("status %d: expected an error, got nil", status)
			}
		})
	}
}



func TestGetRequest_InvalidUrl_ReturnsError(t *testing.T) {
	_, err := GetRequest("http://127.0.0.1:0/nothing", 1)
	if err == nil {
		t.Error("expected error for unreachable URL, got nil")
	}
}

func TestGetRequest_EmptyBody_ReturnsEmptySlice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	body, err := GetRequest(server.URL, 5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(body) != 0 {
		t.Errorf("expected empty body, got %q", string(body))
	}
}