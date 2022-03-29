package router

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type router struct {
	Port   string
	Debug  bool
	Engine *gin.Engine
}

type Router interface {
	SetDebug(opt bool)
	SetPort(pn int) error
	Invoke()
	Run() error
}

func NewRouter() Router {
	return &router{
		Port: ":8050",
	}
}

func (r *router) SetDebug(opt bool) { r.Debug = opt }

func (r *router) SetPort(pn int) error {
	if pn < 0 || pn > 65535 {
		return errors.New("input port number is out of range")
	}
	r.Port = fmt.Sprintf(":%d", pn)
	return nil
}

func (r *router) Invoke() {
	if !r.Debug {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	g := gin.Default()
	g.SetTrustedProxies(nil)

	r.Engine = g
}

func (r *router) Run() error {
	if r.Port == "" {
		// switch to default port
		r.Port = ":8050"
	}
	if r.Engine == nil {
		return errors.New("gin engine has not been invoked yet")
	}

	return r.Engine.Run(r.Port)
}
