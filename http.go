package tools

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HttpValues map[string]string

// NewBytesResponseHTTP return stream response as slice of bytes (standard behavior)
func NewBytesResponseHTTP(res *http.Response) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, res.ContentLength))
	_, err := buf.ReadFrom(res.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERR ] Could not read Response's Body: %s", err)
	}
	if res.StatusCode != 200 {
		log.Printf("[ERR ] Error in response, StatusCode=%d, reason:%s", res.StatusCode, buf.String())
	}
	return buf.Bytes(), nil
}

// NewStringResponseHTTP allow to transform a stream response into string
func NewStringResponseHTTP(res *http.Response) (string, error) {
	b, err := NewBytesResponseHTTP(res)
	return string(b), err
}

func makeRequest(body, headers map[string]string, endpoint, method string) *http.Request {
	var dataSerialized io.Reader
	data := url.Values{}

	for k, v := range body {
		data.Set(k, v)
	}

	if len(data) > 0 {
		dataSerialized = strings.NewReader(data.Encode())
	}

	req, err := http.NewRequest(method, endpoint, dataSerialized)

	if err != nil {
		log.Printf("[ERR ] Could not make request Client: %s", err)
		return nil
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req
}

func sendSimpleGetRequest(body, headers HttpValues, endpoint string) (*http.Response, error) {
	client := &http.Client{}
	req := makeRequest(body, headers, endpoint, "GET")
	return client.Do(req)
}
