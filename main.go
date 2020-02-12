package main

import (
	"context"
	"github.com/AnhNguyenQuoc/Kubernetes-ready/handlers"
	"github.com/AnhNguyenQuoc/Kubernetes-ready/version"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {


	log.Printf("Starting the service ...\n commit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}
	r := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	svr := &http.Server{
		Addr: ":" + port,
		Handler: r,
	}

	go func() {
		log.Fatal(svr.ListenAndServe())
	}()


	log.Print("The service is ready to listen and serve.")
	killSignal := <- interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("GOT SIGTERM...")
	}

	log.Print("The service is shutting down...")
	svr.Shutdown(context.Background())
	log.Print("Done")
}
