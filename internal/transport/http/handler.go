package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRouter() {
	fmt.Println("here we go")

	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "I m alive")

	})
}
