package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Client is an interface meant to wrap around net/http package
// for mocking purpose.
//
// Must only contains a Do() method
type Client interface {
	Do(Body, Values, string, string) (*http.Response, error)
}

// Request simply wraps around net/http
type Default struct {
	Client *http.Client
}

// DefaultClient is used by every helper function
// as a base client
var DefaultClient = &Default{
	Client: &http.Client{
		Timeout: time.Second * 10,
	},
}

// Values is meant to stock Headers and Body values
type Values map[string]string

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
func BuildRequest(body Body, headers Values, endpoint, method string) *http.Request {
	var b io.Reader

	if body != nil {
		b = body.GetBody()
	}

	req, err := http.NewRequest(method, endpoint, b)

	if err != nil {
		log.Printf("Could not make request Client: %s", err)
		return nil
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req
}

// Request is a helper using DefaultClient
func Request(body Body, headers Values, endpoint, method string) (*http.Response, error) {
	return DefaultClient.Do(body, headers, endpoint, method)
}

// Do implements Client interface
func (r *Default) Do(body Body, headers Values, endpoint, method string) (*http.Response, error) {
	req := BuildRequest(body, headers, endpoint, method)
	return r.Client.Do(req)
}

// SendXWWWFormURLEncodedRequest sends a x-www-form-url-encoded thru POST method
func SendXWWWFormURLEncodedRequest(body Body, headers map[string]string, endpoint string) (*http.Response, error) {
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return DefaultClient.Do(body, headers, endpoint, "POST")
}

// SendSimpleGetRequest allow to send get request to an url
func SendSimpleGetRequest(body Body, headers map[string]string, endpoint string) (*http.Response, error) {
	return DefaultClient.Do(body, headers, endpoint, "GET")
}
