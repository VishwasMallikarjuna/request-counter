package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/VishwasMallikarjuna/request-counter/core/counter"
)

// Handler is responsible for managing and responding to HTTP requests.
type Handler struct {
	regAccess counter.RegisterAccess
}

// NewHandler creates and returns a new Handler instance.
func NewHandler(regAccess counter.RegisterAccess) Handler {
	return Handler{regAccess: regAccess}
}

// ServeHTTP processes incoming HTTP requests and responds with the request count
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.regAccess.Access(time.Now().Unix())

	count := h.regAccess.Register()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(count); err != nil {
		log.Printf("Error encoding response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
