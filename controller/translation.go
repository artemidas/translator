package controller

import (
	"encoding/json"
	"github.com/artemidas/translator/model"
	"github.com/artemidas/translator/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type TranslationController struct {
	db *mongo.Client
}

func NewTranslationController(db *mongo.Client) *TranslationController {
	return &TranslationController{
		db: db,
	}
}

func (tc *TranslationController) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello World!"})
}

func (tc *TranslationController) Translations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("getting translations for " + vars["lang"]))
}

func (tc *TranslationController) CreateTranslation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := model.Translation{}
	if err := utils.DecodeBody(r, &t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := t.Insert(tc.db, vars["lang"])
	if err != nil {
		log.Fatal("Error creating translation: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func (tc *TranslationController) UpdateTranslation(w http.ResponseWriter, r *http.Request) {

}
