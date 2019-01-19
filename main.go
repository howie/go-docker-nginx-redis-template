package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/howie/go11x-gae-project-template/internal"
	"google.golang.org/appengine"
	"net/http"
	"os"
)

func main() {

	registerHandlers()
	// For Go1.9 runtime
	appengine.Main()
}

func registerHandlers() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/echo", Echo).Methods("POST")

	// For Go1.9 runtime
	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
	// [END request_logging]

	// For Go1.11 runtime
	//  port := os.Getenv("PORT")
	//  if port == "" {
	//	port = "8080"
	//	log.Printf("Defaulting to port %s", port)
	//  }
	//  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
