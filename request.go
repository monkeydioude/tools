package tools

import "net/http"

func SendXWWWFormUrlEncodedRequest(body, headers HttpValues, endpoint string) (*http.Response, error) {
	client := &http.Client{}
	req := makeRequest(body, headers, endpoint, "POST")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}

func SendSimpleGetRequest(body, headers HttpValues, endpoint string) (*http.Response, error) {
	client := &http.Client{}
	req := makeRequest(body, headers, endpoint, "GET")
	return client.Do(req)
}
