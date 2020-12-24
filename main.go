package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/product-api/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	ph := handlers.NewLoggingProducts(l)

	// main handler
	sm := mux.NewRouter().StrictSlash(true)

	router := sm.PathPrefix("/api/v1").Subrouter()

	popRouter := router.Methods(http.MethodPost).Subrouter()
	popRouter.HandleFunc("/products/initial_data", ph.PopulateDataLogging)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetLoggingProducts)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.PostLoggingProducts)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{[0-9]+}", ph.DeleteLoggingProducts)

	updateRouter := router.Methods(http.MethodPost).Subrouter()
	updateRouter.HandleFunc("/products/{[0-9]+}", ph.UpdateLoggingProducts)

	s := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

}
