package rates

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchRates(t *testing.T) {
	mockResponse := `{
		"asks": [["1.12", "100"]],
		"bids": [["1.10", "200"]]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	rate, err := FetchRates(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if rate.Ask != 1.12 {
		t.Errorf("expected ask=1.12, got %v", rate.Ask)
	}
	if rate.Bid != 1.10 {
		t.Errorf("expected bid=1.10, got %v", rate.Bid)
	}
}
