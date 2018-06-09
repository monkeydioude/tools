package tools

import (
	"fmt"
	"net/http"
)

const (
	notFound = "Not Found"
)

type ResponseError struct {
	msg string
}

func (r ResponseError) String() string {
	return ""
}

func Response404(err error) ([]byte, int, error) {
	return nil, 404, err
}

func Response500(err error) ([]byte, int, error) {
	return nil, 500, err
}

func Response200(data []byte) ([]byte, int, error) {
	return data, 200, nil
}

func HttpResponse(rw http.ResponseWriter, msg string, code int) {
	http.Error(rw, fmt.Sprintf(`{"code": %d, "message": "%s"}`, code, msg), code)
}

func HttpNotFound(rw http.ResponseWriter) {
	HttpResponse(rw, notFound, 404)
}
