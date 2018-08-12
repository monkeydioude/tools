package rjson

import (
	"encoding/json"

	"github.com/monkeydioude/tools/http"
)

// Request sends a request, read response body and json.Unmarshal into the passed entity
func Request(endpoint string, method string, body http.Body, headers http.HttpValues, entity interface{}) (interface{}, error) {
	res, err := http.Request(body, headers, endpoint, "GET")
	if err != nil {
		return nil, err
	}

	bytes, err := http.NewBytesResponseHTTP(res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Get calls RequestJsonEntity forcing a GET method
func Get(endpoint string, body http.Body, headers http.HttpValues, entity interface{}) (interface{}, error) {
	return Request(endpoint, "GET", body, headers, entity)
}

// Post calls RequestJsonEntity forcing a POST method
func Post(endpoint string, body http.Body, headers http.HttpValues, entity interface{}) (interface{}, error) {
	return Request(endpoint, "POST", body, headers, entity)
}

func PostEntity(endpoint string, body interface{}, headers http.HttpValues, entity interface{}) (interface{}, error) {
	return Post(endpoint, http.NewJSONEntityBody(body), headers, entity)
}
