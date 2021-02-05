package controller

import (
	"encoding/json"
	"github.com/artemidas/translator/database"
	"github.com/artemidas/translator/model"
	"github.com/artemidas/translator/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello World!"})
}

func Translations(w http.ResponseWriter, r *http.Request) {
	db := database.NewMongo()
	vars := mux.Vars(r)
	t := model.Translation{}
	translations, err := t.GetLocale(db, vars["lang"])
	if err != nil {
		utils.RespondHttpError(w, http.StatusInternalServerError, "Error getting ")
		return
	}
	utils.RespondJson(w, http.StatusOK, &translations)
}

func CreateTranslation(w http.ResponseWriter, r *http.Request) {
	db := database.NewMongo()
	vars := mux.Vars(r)
	t := model.Translation{}
	if err := utils.DecodeBody(r, &t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := t.Insert(db, vars["lang"])
	if err != nil {
		log.Fatal("Error creating translation: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateTranslation(w http.ResponseWriter, r *http.Request) {

}
