package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/url"
	"strings"
)

// Body defines the interface taken by MakeRequest
type Body interface {
	GetBody() io.Reader
}

// JSONEntityBody implements Body interface
// Is used to pass to MakeRequest any JSON entity
type JSONEntityBody struct {
	body io.Reader
}

// NewJSONEntityBody generates pointer to JSONEntityBody from interface
// Entity should be meant to be converted into json through Marshal
func NewJSONEntityBody(entity interface{}) *JSONEntityBody {
	res, err := json.Marshal(entity)

	if err != nil {
		log.Printf("[ERR ] Could not unmarshal entity. Reason: %s\n", err)
		return nil
	}

	return &JSONEntityBody{
		body: bytes.NewReader(res),
	}
}

// GetBody implements Body interface
func (b *JSONEntityBody) GetBody() io.Reader {
	return b.body
}

// ValuesBody implments Body interface
// Is used to pass to MakeRequest any HttpValues body
type ValuesBody struct {
	body io.Reader
}

// NewValuesBody generates pointer to ValuesBody from HttpValues var
func NewValuesBody(body HttpValues) *ValuesBody {
	data := url.Values{}

	for k, v := range body {
		data.Set(k, v)
	}

	if len(data) <= 0 {
		log.Println("[ERR ] Could not create Body, len(data) negative or 0")
		return nil
	}
	return &ValuesBody{
		body: strings.NewReader(data.Encode()),
	}
}

// GetBody implements Body interface
func (b *ValuesBody) GetBody() io.Reader {
	return b.body
}
