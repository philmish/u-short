package internal

import (
	"fmt"
	"net/http"

	"github.com/philmish/u-short/internal/middleware"
)

type ServerConf struct {
    Addr string
    Handler *RegexMatcher
    Middleware []middleware.Middleware
}

func (c ServerConf)DefaultServer() *http.Server {
    c.Handler.Add("/echo/[a-zA-Z0-9]{0,14}", middleware.ChainMiddleware(echo, c.Middleware...))
    return &http.Server{
        Addr: c.Addr,
        Handler: c.Handler,
    }
}

func commonMiddleware() []middleware.Middleware {
    return []middleware.Middleware{middleware.Logger}
}

func DevServer(port int) *http.Server {
    ware := commonMiddleware()
    handler := NewRegexMatcher()
    conf := ServerConf{
        fmt.Sprintf(":%d", port),
        handler,
        ware,
    }
    return conf.DefaultServer()
}
