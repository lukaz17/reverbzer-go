// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tforce-io/tf-golib/opx"
)

// Handles all incoming HTTP requests
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to handle OPTIONS requests properly
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle OPTIONS preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reverbzer: OK"))
		return
	}

	// Send request detail to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	requestInfo := NewRequestInfo(r)
	response := opx.Must1(json.Marshal(requestInfo))
	w.Write(response)
}

// Start the HTTP server
func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RequestHandler)
	server := &http.Server{
		Addr:    ":11111",
		Handler: mux,
	}

	fmt.Println("Starting web server on port 11111...")
	fmt.Println("The server will log all incoming requests (URI, headers, body)")
	fmt.Println("Supported methods: GET, POST, PUT, DELETE, OPTIONS")
	fmt.Println("Press Ctrl+C to stop the server")
	fmt.Println(strings.Repeat("=", 51))

	log.Fatal(server.ListenAndServe())
}
