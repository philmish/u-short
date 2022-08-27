package internal

import (
    "net/http"
    "regexp"
    "fmt"
)

type RegexMatcher struct {
    Patterns map[string]*regexp.Regexp
    Handlers map[string]http.HandlerFunc
}

func NewRegexMatcher() *RegexMatcher {
    return &RegexMatcher{
        Patterns: make(map[string]*regexp.Regexp),
        Handlers: make(map[string]http.HandlerFunc),
    }
}

func (r *RegexMatcher)Add(pattern string, h http.HandlerFunc) error {
    compiled, err := regexp.Compile(pattern)
    if err != nil {
        return fmt.Errorf("Could not compile regex: %s", err.Error())
    }
    r.Handlers[pattern] = h
    r.Patterns[pattern] = compiled

    return nil
}

func (r *RegexMatcher)ServeHTTP(res http.ResponseWriter, req *http.Request) {
    toMatch := req.Method + " " + req.URL.Path
    for regex, handler := range r.Handlers {
        if r.Patterns[regex].MatchString(toMatch) {
            handler(res, req)
            return
        }
    }
    http.NotFound(res, req)
}
