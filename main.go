// Simple web server in Go
// All go applications start with package main, which defines a standalone executable program
package main

// Import necessary packages
import (
	"log"      // Package for logging
	"net/http" // Package for HTTP client and server implementations
	"time"     // Package for time-related functions

	"github.com/gorilla/mux"
)

func getCurrentTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := getCurrentTime()
	w.Write([]byte("Current Time: " + currentTime))
}

// Main function - entry point of the application
func main() {
	// Set up HTTP router
	router := mux.NewRouter()

	// register getTimeHandler
	router.HandleFunc("/time", getTimeHandler).Methods("GET")

	log.Println("starting server on port 8080")
	//fmt.Println("Current Time:", getCurrentTime())
	// Start the HTTP server on port 8080

	http.ListenAndServe(":8080", router)

}
