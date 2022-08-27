package internal

import (
	"fmt"
	"net/http"
)

type ServerConf struct {
    Addr string
    Handler http.Handler
}

func (c ServerConf)DefaultServer() error {
    return http.ListenAndServe(c.Addr, c.Handler)
}

func DevServer(port int) error {
    conf := ServerConf{fmt.Sprintf(":%d", port), nil}
    return conf.DefaultServer()
}
