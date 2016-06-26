package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/pat"
	"github.com/russross/blackfriday"
)

//Post is a blog post
type Post struct {
	Body string    `json:"body"`
	Time time.Time `json:"time"`
}

var db []Post

func main() {
	fmt.Println("Hello, World!")

	r := pat.New()

	r.Get("/hello", hello)
	r.Post("/markdown", markdown)
	r.Post("/posts", addPost)
	r.Get("/posts", getPosts)

	r.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on localhost:", port)
	http.ListenAndServe(":"+port, r)

}

func hello(w http.ResponseWriter, r *http.Request) {
	body := []byte("Hello, World!")

	w.Write(body)

}

func markdown(w http.ResponseWriter, r *http.Request) {
	body := []byte(r.FormValue("body"))
	markdown := blackfriday.MarkdownCommon(body)
	w.Write(markdown)

}

//responsible for adding a new post
func addPost(w http.ResponseWriter, r *http.Request) {
	var p Post

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Time = time.Now()

	db = append(db, p)

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Print(err)
	}

}

func getPosts(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
