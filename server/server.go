// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rs/zerolog"
)

// Start the HTTP server
func Run(port uint16, logger zerolog.Logger) {
	mux := http.NewServeMux()
	handler := &Handler{
		Logger: logger,
	}
	mux.HandleFunc("/", handler.RequestHandler)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("Starting web server on port %d...\n", port)
	fmt.Println("The server will log all incoming requests (URI, headers, body)")
	fmt.Println("Supported methods: GET, POST, PUT, DELETE, OPTIONS")
	fmt.Println("Press Ctrl+C to stop the server")
	fmt.Println(strings.Repeat("=", 51))

	log.Fatal(server.ListenAndServe())
}
