package rjson

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/monkeydioude/tools"
)

// RequestEntity sends a request, read response body and json.Unmarshal into the passed entity
func RequestEntity(method string, body, headers tools.HttpValues, endpoint string, entity interface{}) (interface{}, error) {
	client := &http.Client{}
	req := tools.MakeRequest(body, headers, endpoint, "GET")
	if req == nil {
		return nil, errors.New("Request was empty")
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := tools.NewBytesResponseHTTP(res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// GetEntity calls RequestJsonEntity forcing a GET method
func GetEntity(body, headers tools.HttpValues, endpoint string, entity interface{}) (interface{}, error) {
	return RequestEntity("GET", body, headers, endpoint, entity)
}
