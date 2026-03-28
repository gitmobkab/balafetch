package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{}

// GetRequest is an utility function to execute a get request on the internet
// a response not in the range 200-299 is flagged as bad and cause it to return an error
func GetRequest(url string, requestTimeout int) ([]byte, error) {

	ctx := context.Background()
	cancel := func() {}
	if requestTimeout != 0 {
		timeoutDuration := time.Duration(requestTimeout)*time.Second
		ctx, cancel = context.WithTimeout(context.Background(), timeoutDuration.Abs())
    }
    defer cancel()


	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}
	
	return io.ReadAll(resp.Body) 
}