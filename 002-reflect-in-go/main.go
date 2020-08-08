package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	/*
		reflection is just a mechanism to examine the type and value pair stored inside an interface variable
	*/
	//TO view reflect.TYPE
	fmt.Println("Type is :", reflect.TypeOf(x))
	//TO view reflect.VALUE (concrete value)
	fmt.Println("Value of x is ::", reflect.ValueOf(x))
	//TO view reflect.VALUE type
	fmt.Println("Value of x is ::", reflect.ValueOf(x).String())
}
