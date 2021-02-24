package controller

import (
	"github.com/artemidas/translator/database"
	"github.com/artemidas/translator/model"
	"github.com/artemidas/translator/utils"
	"github.com/gin-gonic/gin"
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

func GenerateTranslation(c *gin.Context) {
	db := database.NewMongo()
	t := model.Translation{}
	translations, err := t.GetLocale(db, c.Param("lang"))
	if err != nil {
		log.Println("error generating translation:", err.Error())
		code := http.StatusBadRequest
		c.JSON(code, utils.Response{Code: code, Message: "error generating translation"})
		return
	}
	c.Status(http.StatusOK)
	c.Header("Content-Type", "application/json")
	c.Writer.Write([]byte(translations))
}

func CreateTranslation(c *gin.Context) {
	db := database.NewMongo()
	t := model.Translation{}
	if err := c.BindJSON(&t); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := t.Insert(db, c.Param("lang"))
	if err != nil {
		log.Println("error creating translation:", err.Error())
		code := http.StatusBadRequest
		c.JSON(code, utils.Response{Code: code, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{Code: http.StatusOK, Message: "translation created"})
}

func UpdateTranslation(c *gin.Context) {
	db := database.NewMongo()
	t := model.Translation{}
	if err := c.BindJSON(&t); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := t.Update(db, c.Param("lang"), c.Param("id"))
	if err != nil {
		log.Println("error updating translation:", err.Error())
		code := http.StatusBadRequest
		c.JSON(code, utils.Response{Code: code, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{Code: http.StatusOK, Message: "Successfully updated"})
}

func DeleteTranslation(c *gin.Context) {
	db := database.NewMongo()
	t := model.Translation{}
	err := t.Delete(db, c.Param("lang"), c.Param("id"))
	if err != nil {
		log.Println("error deleting translation:", err.Error())
		code := http.StatusBadRequest
		c.JSON(code, utils.Response{Code: code, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{Code: http.StatusOK, Message: "successfully deleted"})

}
