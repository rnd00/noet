package router

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// router

type routern struct {
	Port    string
	Timeout time.Duration
	Engine  *http.Server
	Handler *http.Handler
}

type Routern interface {
	SetHandler(h *http.Handler) error
	SetPort(pn int) error
	SetTimeout(duration time.Duration) error
	Invoke() error
	Run() error
}

func NewRoutern() Routern {
	return &routern{
		Port: ":8050",
	}
}

func (r *routern) SetHandler(h *http.Handler) error {
	if h == nil {
		return errors.New("parameter is empty")
	}

	r.Handler = h
	return nil
}

func (r *routern) SetPort(pn int) error {
	if pn < 0 || pn > 65535 {
		return errors.New("input port number is out of range")
	}
	r.Port = fmt.Sprintf(":%d", pn)
	return nil
}

func (r *routern) SetTimeout(duration time.Duration) error {
	if duration < 100*time.Millisecond || duration > 30*time.Second {
		return errors.New("duration can't be smaller than 100ms or larger than 30s")
	}
	r.Timeout = duration

	return nil
}

func (r *routern) Invoke() error {
	if r.Handler == nil {
		return errors.New("please set handler first")
	}
	if r.Timeout < 100*time.Millisecond {
		// set default to 5s
		r.Timeout = 5 * time.Second
	}
	newServer := &http.Server{
		Addr:        r.Port,
		Handler:     *r.Handler,
		ReadTimeout: r.Timeout,
	}

	r.Engine = newServer
	return nil
}

func (r *routern) Run() error {
	return r.Engine.ListenAndServe()
}

// handler related func
// (local struct) handler -> (interface) http.Handler

type handler struct {
	Mux map[string]func(http.ResponseWriter, *http.Request)
}

func NewHandler() *handler {
	newMux := make(map[string]func(http.ResponseWriter, *http.Request))
	newHandler := &handler{
		Mux: newMux,
	}
	return newHandler
}

func (h *handler) ReturnHttpHandler() http.Handler {
	return h
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := h.Mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL: "+r.URL.String())
}

func (h *handler) SetupMuxer(path string, function func(http.ResponseWriter, *http.Request)) error {
	if path == "" {
		return errors.New("invalid path, path is empty")
	}
	if path[0] != '/' {
		return errors.New("invalid path, need to begin with / (forward slash)")
	}

	if function == nil {
		return errors.New("invalid function, function parameter is empty")
	}

	h.Mux[path] = function
	return nil
}
