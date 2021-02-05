package main

import (
	"github.com/artemidas/translator/controller"
	"github.com/artemidas/translator/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Live server refresh
// https://github.com/cosmtrek/air

func main() {
	router := mux.NewRouter().StrictSlash(true)

	db := database.NewMongo()

	translationController := controller.NewTranslationController(db)

	router.HandleFunc("/", translationController.Home)
	translator := router.PathPrefix("/translate").Subrouter()

	translator.
		HandleFunc("/{lang}", translationController.Translations).
		Methods(http.MethodGet)

	translator.
		HandleFunc("/{lang}/create", translationController.CreateTranslation).
		Methods(http.MethodPost)

	translator.
		HandleFunc("/{lang}/update/{id}", translationController.UpdateTranslation).
		Methods(http.MethodPut)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
