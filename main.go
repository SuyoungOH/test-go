package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test-go/routers"
	"time"
)

func main() {
	fmt.Println("Hello Web Server")

	mux := routers.ApiRouter()

	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: mux,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error %s\n", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	fmt.Println("HTTP server is closing...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("HTTP server shutdown error %s\n", err)
	}

	fmt.Println("HTTP server is graceful shutdown")

	// 서버가 완전히 종료될때까지 대기
	<-ctx.Done()
	fmt.Println("HTTP server is closed")
}
