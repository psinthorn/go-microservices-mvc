package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/psinthorn/go-microservices-mvc/controllers"
)

func StartApp() {

	os.Setenv("PORT", "8008")
	port := os.Getenv("PORT")

	http.HandleFunc("/", controllers.Welcome)
	http.HandleFunc("/users", controllers.GetUser)
	http.HandleFunc("/todos", controllers.TodoIndex)
	//http.ListenAndServe(":8080", nil)

	fmt.Println("Server running on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}
}
