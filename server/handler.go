// Copyright (C) Nguyen Nhat Tung
//
// Reverbzer is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package server

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/tforce-io/tf-golib/opx"
)

type Handler struct {
	Logger zerolog.Logger
}

// Handles all incoming HTTP requests
func (h *Handler) RequestHandler(w http.ResponseWriter, r *http.Request) {
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

	h.Logger.Info().
		Msgf("%s %s", r.Method, r.RequestURI)
}
