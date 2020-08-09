package models

import "gopkg.in/mgo.v2/bson"

type Dish struct {
	Name           string
	Price          int
	Cuisine        string
	ChefSpeciality bool
	Id             bson.ObjectId
}
