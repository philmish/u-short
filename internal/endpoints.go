package internal

import (
	"fmt"
	"net/http"
	"strings"
)

func echo(res http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    path := req.URL.Path
    parts := strings.Split(path, "/")

    msg := "No Input"
    if len(parts) > 2 {
        msg = parts[2]
    }

    payload := fmt.Sprintf("msg: %s\n", msg)
    res.Write([]byte(payload))
}


