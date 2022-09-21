package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/philmish/s-tree/kvdb"
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
	client := kvdb.DBClient{Addr: "/tmp/ushort"}
	path := req.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) > 2 {
		key := parts[2]
		uri, err := client.Get(key)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(res, req, uri, 302)
		return
	} else {
		http.NotFound(res, req)
		return
	}
}

type shortenReq struct {
	Url string `json:"url"`
}

func (sr shortenReq) validate() bool {
	_, err := url.ParseRequestURI(sr.Url)
	return err == nil
}

func (sr shortenReq) short(taken strslice) string {
	return shorten(sr.Url, taken)
}

func shortenUrl(res http.ResponseWriter, req *http.Request) {
	client := kvdb.DBClient{Addr: "/tmp/ushort"}
	defer req.Body.Close()
	var request shortenReq

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(res, "Failed to decode request", http.StatusBadRequest)
		return
	}
	if !request.validate() {
		http.Error(res, "Invalid url", http.StatusBadRequest)
		return
	}
	taken, err := client.Keys()
	if err != nil {
		http.Error(res, "Something went wrong checking the keys", http.StatusInternalServerError)
		return
	}
	key := request.short(taken)
	err = client.Set(key, request.Url)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Write([]byte(key))
}
