package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/navendu-pottekkat/students-api-test/handlers"
)

var port = ":9090"

func main() {
	l := log.New(os.Stdout, "students-api-test ", log.LstdFlags)
	sh := handlers.NewStudents(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", sh.GetStudents)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", sh.AddStudent)
	postRouter.Use(sh.MiddlewareValidateStudent)

	// Handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sph := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sph)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Create a new server
	s := http.Server{
		Addr:         port,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Printf("Starting server on port %s", port)

		err := s.ListenAndServe()

		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// Gracefully shutdown the server, waiting a max of 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
