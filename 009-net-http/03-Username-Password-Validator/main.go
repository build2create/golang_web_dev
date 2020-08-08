package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"unicode"
)

var tpl *template.Template
var count = 0

type validator struct {
	err bool
	msg string
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}
func (v validator) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	if count > 0 {
		password := req.FormValue("password")
		s := validate(password)
		fmt.Println(s)
		tpl.ExecuteTemplate(w, "template.gohtml", s.msg)
	} else {
		tpl.ExecuteTemplate(w, "template.gohtml", nil)
	}
	fmt.Println("count::", count)
	count++
}
func validate(p string) validator {
	var lengthValid bool
	var specialPresent bool
	var numberPresent bool
	var upperPresent bool
	var lowerPresent bool
	var totalLen int
	var errString string
	var error bool
	for _, c := range p {
		switch {
		case unicode.IsNumber(c):
			numberPresent = true
			totalLen++
		case unicode.IsUpper(c):
			upperPresent = true
			totalLen++
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			specialPresent = true
			totalLen++
		case unicode.IsLower(c):
			lowerPresent = true
			totalLen++
		}
	}
	if totalLen > 8 || totalLen < 64 {
		lengthValid = true
	}

	appendError := func(err string) {
		if len(strings.TrimSpace(errString)) > 0 {
			errString += err
		} else {
			errString = err
		}
	}
	if !lengthValid {
		appendError("Error: Please ensure that the password length is between 8 to 64 characters\n")
		error = true
	}
	if !lowerPresent {
		appendError("Error: Please ensure that at least one lower case character is present\n")
		error = true
	}
	if !upperPresent {
		appendError("Error: Please ensure that at least one upper case character is present\n")
		error = true
	}
	if !numberPresent {
		appendError("Error: Please ensure that there is at least one number is present\n")
		error = true
	}
	if !specialPresent {
		appendError("Error: Please ensure that at least one special character is present\n")
		error = true
	}
	if specialPresent && numberPresent && upperPresent && lowerPresent && lengthValid {
		appendError("Success")
		error = false
	}
	return validator{error, errString}
}
func main() {
	var v validator
	http.ListenAndServe(":8080", v)
}
