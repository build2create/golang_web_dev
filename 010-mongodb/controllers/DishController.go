package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/build2create/golang_web_dev/010-mongodb/models"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DishController is a type to controller the flow between the database and the views
type DishController struct {
	session *mgo.Session
}

// NewDishController returns new dish controller pointer
func NewDishController(s *mgo.Session) *DishController {
	return &DishController{s}
}

//GetDish : http://localhost:3000/dish/<dishName> will return the dish with name dishName
func (dc DishController) GetDish(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//ObjectIdHex gives the ObjectId from the hex representation of the string
	oid := bson.ObjectIdHex(id)

	d := models.Dish{}
	//Fetch the dish
	if err := dc.session.DB("bozzopub").C("dishes").FindId(oid).One(&d); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	dj, _ := json.Marshal(d)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", dj)
}

// CreateDish : will create a dish with params in the POST request
// curl -X POST -H "Content-Type: application/json" -d '{"Name":"Idli","Price":100, "Cuisine":"South","ChefSpeciality":true}' http://localhost:3000/dish
func (dc DishController) CreateDish(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	d := models.Dish{}
	json.NewDecoder(r.Body).Decode(&d)
	//Create a new Id
	d.Id = bson.NewObjectId()
	//Insert document in a sepecific collection
	dc.session.DB("bozzopub").C("dishes").Insert(d)

	dj, _ := json.Marshal(d)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", dj)
}

//DeleteDish : will write code to delete the dish
//curl -X DELETE -H "Content-Type: application/json" http://localhost:3000/dish/DOSA
func (dc DishController) DeleteDish(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//ObjectIdHex gives the ObjectId from the hex representation of the string
	oid := bson.ObjectIdHex(id)

	//Delete the user
	if err := dc.session.DB("bozzopub").C("dishes").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s\n", "Write some code to delete the dish")
}
