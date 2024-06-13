package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {

	//connect to database
	s, err := mgo.Dial("mongodb://localhost:27107")

	if err != nil {
		panic(err)
	}
	return s
}

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	//routes
	//To do: Update, Get All
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// start server
	http.ListenAndServe("localhost:9000", r)

}
