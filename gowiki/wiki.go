package main

import (
//	"fmt"
	"log"
	"net/http"
	
	"Handler"
	"Edit"
	"View"
	"Save"
)

/*
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m != nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}*/

func main () {
	http.HandleFunc("/view/", Handler.MakeHandler(View.ViewHandler))
	http.HandleFunc("/edit/", Handler.MakeHandler(Edit.EditHandler))
	http.HandleFunc("/save/", Handler.MakeHandler(Save.SaveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}



