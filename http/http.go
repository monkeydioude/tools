package http

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// Values is a helper to build simple Headers
type HttpValues map[string]string

// NewBytesResponseHTTP return stream response as slice of bytes (standard behavior)
func NewBytesResponseHTTP(res *http.Response) ([]byte, error) {
	if res.ContentLength <= 0 {
		res.ContentLength = 255
	}
	buf := bytes.NewBuffer(make([]byte, 0, res.ContentLength))
	_, err := buf.ReadFrom(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read Response's Body: %s", err)
	}
	if res.StatusCode != 200 {
		log.Printf("Error in response, StatusCode=%d, reason:%s", res.StatusCode, buf.String())
	}
	return buf.Bytes(), nil
}

// NewStringResponseHTTP allow to transform a stream response into string
func NewStringResponseHTTP(res *http.Response) (string, error) {
	b, err := NewBytesResponseHTTP(res)
	return string(b), err
}

// BuildRequest allows to make custom request using body, headers, url and a method
// Most likely wrapped by more practical functions (cf: SendXWWWFormUrlEncodedRequest)
func BuildRequest(body Body, headers map[string]string, endpoint, method string) *http.Request {
	req, err := http.NewRequest(method, endpoint, body.GetBody())

	if err != nil {
		log.Printf("Could not make request Client: %s", err)
		return nil
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req
}

func Request(body Body, headers map[string]string, endpoint, method string) (*http.Response, error) {
	client := &http.Client{}
	req := BuildRequest(body, headers, endpoint, method)
	return client.Do(req)
}

// SendXWWWFormUrlEncodedRequest sends a x-www-form-url-encoded thru POST method
func SendXWWWFormUrlEncodedRequest(body Body, headers map[string]string, endpoint string) (*http.Response, error) {
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return Request(body, headers, endpoint, "POST")
}

// SendSimpleGetRequest allow to send get request to an url
func SendSimpleGetRequest(body Body, headers map[string]string, endpoint string) (*http.Response, error) {
	return Request(body, headers, endpoint, "GET")
}
