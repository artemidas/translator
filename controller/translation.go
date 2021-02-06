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

func GenerateTranslation(w http.ResponseWriter, r *http.Request) {
	db := database.NewMongo()
	vars := mux.Vars(r)
	t := model.Translation{}
	translations, err := t.GetLocale(db, vars["lang"])
	if err != nil {
		log.Println("error generating translation:", err.Error())
		utils.RespondHttpError(w, http.StatusInternalServerError, "error generating translation")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(translations))
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
		log.Println("error creating translation:", err.Error())
		utils.RespondHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateTranslation(w http.ResponseWriter, r *http.Request) {
	db := database.NewMongo()
	vars := mux.Vars(r)
	t := model.Translation{}
	if err := utils.DecodeBody(r, &t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := t.Update(db, vars["lang"], vars["id"])
	if err != nil {
		log.Println("error updating translation:", err.Error())
		utils.RespondHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJson(w, http.StatusOK, &utils.Response{Code: http.StatusOK, Message: "Successfully updated"})
}
