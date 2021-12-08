package main

import (
	"context"
	"fmt"
	"getir/app"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed loading .env file")
	}
}
func main() {
	appModule := new(app.App)
	appModule.Initialize()

	http.HandleFunc("/", appModule.Handler)
	http.HandleFunc("/records", appModule.RecordListHandler)
	http.HandleFunc("/in-memory", appModule.InMemoryHandler)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	httpServer := &http.Server{
		Addr: port,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	log.Println("Server Listening on port", port)

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan
	log.Println("Server Stopped")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Println("gracefully stopped")
	}

	defer os.Exit(0)
}
