package main

import (
	"crypto/tls"
	"io"
	"net/http"
)

// FetchInsecure uses an HTTP client that disables TLS verification
func FetchInsecure(url string) ([]byte, error) {
	// VULNERABLE: InsecureSkipVerify = true disables certificate validation.
	// Static analysers typically flag this as an insecure TLS configuration.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
