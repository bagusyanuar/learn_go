package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"web_go/database"
)

func MainPage(response http.ResponseWriter, request *http.Request) {
	// var filePath = path.Join("views", "index.html")
	t, err := template.ParseFiles("views/layouts.html", "views/index.html")

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	err = t.ExecuteTemplate(response, "base", nil)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func AboutPage(response http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("views/layouts.html", "views/about.html")

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	err = t.ExecuteTemplate(response, "base", nil)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func SampleApi(response http.ResponseWriter, request *http.Request) {
	var cities []City
	database.Db.Find(&cities)
	jsonInBytes, err := json.Marshal(cities)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(cities)
	response.Write(jsonInBytes)
}

type City struct {
	Id   string
	Name string
	Slug string
}

func (City) TableName() string {
	return "cities"
}
