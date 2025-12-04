package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shariyerShazan/GoLang-Crud-MongoDB/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// 2) Router
	router := httprouter.New()

	// 3) User Controller instance
	userController := controllers.NewUserController(getSession())

	// 4) Routes
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user", userController.CreateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	// 5) Server start
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}


func getSession() *mgo.Session{
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Mongo connect error:", err)
		panic(err)
	}
	return session
}