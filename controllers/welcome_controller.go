package controllers

import (
	"encoding/json"
	"net/http"
)

func Welcome(res http.ResponseWriter, req *http.Request) {
	contentData, _ := json.Marshal(map[string]string{"title": "Welcome to Microservices and MVC by Golang"})

	res.WriteHeader(http.StatusOK)
	res.Write(contentData)
}
