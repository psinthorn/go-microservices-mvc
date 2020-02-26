package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/psinthorn/go-microservices-mvc/services"
	"github.com/psinthorn/go-microservices-mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userIdErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(userIdErr)
		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonValue)
		return
	}

	user, userIdErr := services.GetUser(userId)

	if userIdErr != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)

}

func EditUser(res http.ResponseWriter, req *http.Request) {
	/* Step and Requirment Preparing
	1. Template
	2. Data
	3. Render
	*/

	//Template
	tmpl := template.Must(template.ParseFiles("./views/users/edit.html"))

	tmpl.Execute(res, "Edit User")

}
