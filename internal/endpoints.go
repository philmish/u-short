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

    msg := ""
    if len(parts) > 2 {
        msg = parts[2]
    } else {
        msg = "Missing Input"
    }

    payload := fmt.Sprintf("msg: %s\n", msg)
    res.Write([]byte(payload))
}

func redirect(res http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()
    //TODO implement storage solution
    mapping := map[string]string{
        "g": "https://google.com/",
    }
    path := req.URL.Path
    parts := strings.Split(path, "/")

    if len(parts) > 2 {
        key := parts[2]
        for k, v := range mapping {
            if k == key {
                http.Redirect(res, req, v, 302)
                return
            }
        }
    } else {
        http.NotFound(res, req)
        return
    }
}


