package tools

import (
	"fmt"
	"net/http"
)

const (
	notFound = "Not Found"
)

func response404(err error) ([]byte, int, error) {
	return nil, 404, err
}

func response500(err error) ([]byte, int, error) {
	return nil, 500, err
}

func response200(data []byte) ([]byte, int, error) {
	return data, 200, nil
}

func httpResponse(rw http.ResponseWriter, msg string, code int) {
	http.Error(rw, fmt.Sprintf(`{"code": %d, "message": "%s"}`, code, msg), code)
}

func httpNotFound(rw http.ResponseWriter) {
	httpResponse(rw, notFound, 404)
}
