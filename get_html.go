package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode >= 400 {
		return "", fmt.Errorf("StatusCode: %v - %s", statusCode, resp.Status)
	}

	// Content-Type might have encoding attribute, separated on ';'
	contentType := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	if contentType != "text/html" {
		return "", fmt.Errorf("Content-Type error: %s", contentType)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
