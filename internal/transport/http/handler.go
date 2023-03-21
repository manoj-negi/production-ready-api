package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CommentService interface{}

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	fmt.Println("setting up Routes")

	h.Router.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I m alive")
	})
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
