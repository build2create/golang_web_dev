package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func getUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/*	d := models.Dish{
			Name:           "PuranPoli",
			Price:          "100",
			Cuisine:        "Maharashtrian",
			ChefSpeciality: true,
		}
		dj, _ = json.Marshal(d)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", dj)*/
}
func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/user", getUser)
	log.Fatal(http.ListenAndServe(":3000", router))
}
