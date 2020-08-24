package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

//Do not forget to go get github.com/google/uuid

func serve() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cookie/", cookie)
	http.HandleFunc("/uuid/", setuuid)
	http.ListenAndServe(":8080", nil)
}

func main() {
	serve()
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, `This Is The Main Page
	Go to /cookie/ or /uuid/`)
	if err != nil {
		log.Panicln(err)
	}
}

func cookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "maincookie",
		Value: "This is the main cookie",
		Path:  "/",
	})
	_, err := fmt.Fprintln(w, "Cookie Has Been Successfully Inserted")
	if err != nil {
		log.Panicln(err)
	}
}

func setuuid(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("uuidCookie")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "uuidCookie",
			Value: uuid.New().String(),
			Path:  "/",
		})
		fmt.Fprintln(w, "Success, Inserted A UUID Cookie")
	} else {
		fmt.Fprintln(w, "Already Inserted Cookie")
	}
}
