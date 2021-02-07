package main

import (
	"github.com/artemidas/translator/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	translator := router.PathPrefix("/api/translate").Subrouter()
	translator.HandleFunc("/{lang}", controller.GenerateTranslation).Methods(http.MethodGet)
	translator.HandleFunc("/{lang}/create", controller.CreateTranslation).Methods(http.MethodPost)
	translator.HandleFunc("/{lang}/update/{id}", controller.UpdateTranslation).Methods(http.MethodPut)

	files := http.FileServer(http.Dir("dist"))
	router.PathPrefix("/css").Handler(http.StripPrefix("/", files))
	router.PathPrefix("/img").Handler(http.StripPrefix("/", files))
	router.PathPrefix("/js").Handler(http.StripPrefix("/", files))
	router.PathPrefix("/favicon.ico").Handler(http.StripPrefix("/", files))
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
