package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetRequest(url string) ([]byte, error) {

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	
	body, err := io.ReadAll(resp.Body) 
	resp.Body.Close()

	if resp.StatusCode < 200 && resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected response status code: %d with data: %s", resp.StatusCode, body)
	}

	if err != nil {
		return nil, err

	}

	return body,nil

}