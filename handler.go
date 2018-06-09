package tools

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Configuration map[string]string
type Routes map[string]*Route

type Handler struct {
	config  *Configuration
	headers map[string]string
	Routes  Routes
}

type Xer func([]string, *Configuration) ([]byte, int, error)

type Route struct {
	Handler func([]string, *Configuration) ([]byte, int, error)
	Method  string
}

func (h *Handler) WithHeader(k, v string) {
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[k] = v
}

func (h *Handler) applyHeaders(rw http.ResponseWriter) {
	for key, value := range h.headers {
		rw.Header().Set(key, value)
	}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}

	h.applyHeaders(rw)

	for p, route := range h.Routes {
		if route.Method != r.Method {
			continue
		}
		v, err := MatchAndFind(p, strings.Trim(r.RequestURI, "/"))

		if err != nil {
			continue
		}

		data, _, err := route.Handler(v, h.config)
		if err != nil {
			log.Println(err)
			HttpNotFound(rw)
			return
		}

		fmt.Fprint(rw, string(data))
		return
	}
	log.Printf("[WARN] '%s' did not match any route\n", r.RequestURI)
	HttpNotFound(rw)
}

func NewHandler(conf *Configuration) *Handler {
	return &Handler{
		config: conf,
		Routes: make(Routes),
	}
}

func (routes *Routes) Add(r, m string, f Xer) {
	(*routes)[r] = &Route{
		Method:  m,
		Handler: f,
	}
}

func (routes *Routes) AddGet(r string, f Xer) {
	routes.Add(r, "GET", f)
}
