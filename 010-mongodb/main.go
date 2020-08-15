package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/build2create/golang_web_dev/010-mongodb/controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	dc := controllers.NewDishController(getSession())
	router.GET("/dish/:id", dc.GetDish)
	router.POST("/dish", dc.CreateDish)
	router.DELETE("/dish/:id", dc.DeleteDish)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb//mongodb:27017")
	if err != nil {
		panic(err)
	}
	return s
}
