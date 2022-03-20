package cmd

import (
	"PortalClient/internal/handlers"

	"github.com/gorilla/mux"
)

func HandleRoutes(router *mux.Router) {
	router.HandleFunc("/api/ideas", handlers.ErrorHandler(handlers.GetIdeasHandler))
	router.HandleFunc("/api/postIdea", handlers.ErrorHandler(handlers.PostIdeaHandler)).Methods("POST")
}
