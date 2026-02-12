package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetRequest(url string, request_timeout int) ([]byte, error) {

	client := http.Client{
		Timeout: time.Duration(request_timeout) * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body) 
	resp.Body.Close()
	
	if err != nil {
		return nil, err

	}

	return body,nil

}