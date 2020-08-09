package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/build2create/golang_web_dev/010-mongodb/controllers"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	dc := controllers.NewDishController()
	router.GET("/dish/:name", dc.GetDish)
	router.POST("/dish", dc.CreateDish)
	router.DELETE("/dish/:name", dc.DeleteDish)
	log.Fatal(http.ListenAndServe(":3000", router))
}
