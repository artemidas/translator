package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}

func Translations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("getting translations for " + vars["lang"]))
}

func CreateTranslation(w http.ResponseWriter, r *http.Request) {

}

func UpdateTranslation(w http.ResponseWriter, r *http.Request) {

}
