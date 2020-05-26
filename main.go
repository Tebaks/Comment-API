package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Comment-API/data"

	"github.com/Comment-API/config"
	"github.com/Comment-API/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config.ReadConfig()

	data.DatabaseConnection()

	l := log.New(os.Stdout, "comments-api ", log.LstdFlags)

	sm := mux.NewRouter()

	ch := handlers.NewComments(l)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ch.GetComments)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ch.AddComment)
	postRouter.Use(ch.MiddlewareValidateComment)

	s := http.Server{
		Addr:         config.C.Server.Address, // configure the bind address
		Handler:      sm,                      // set the default handler
		ErrorLog:     l,                       // set the logger for the server
		ReadTimeout:  5 * time.Second,         // max time to read request from the client
		WriteTimeout: 10 * time.Second,        // max time to write response to the client
		IdleTimeout:  120 * time.Second,       // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server...")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
