package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/codertjay/mongo-golang/controllers"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017/")
	if err != nil {
		panic(err)
	}
	return s
}
