package controller

import (
	"github.com/artemidas/translator/database"
	"github.com/artemidas/translator/model"
	"github.com/artemidas/translator/utils"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	index, _ := template.ParseFiles("dist/index.html")
	err := index.Execute(w, nil)
	if err != nil {
		log.Printf("error occurred while executing the template or writing its output: %s", err)
		return
	}
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
