package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/philmish/u-short/internal"
)

func main() {
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    srv := internal.DevServer(9999)

    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
           log.Fatalf("%s\n", err.Error()) 
        }
    }()
    log.Println("Server started")
    <-done

    log.Println("Server stopped")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer func() {
        cancel()
    }()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server shutdown failed: %s", err.Error())
    }
    log.Println("Server shutdown gracefully")
}
