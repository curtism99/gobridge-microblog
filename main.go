package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/pat"
	"github.com/russross/blackfriday"
)

//Post is a blog post
type Post struct {
	Body  Markdown  `json:"body"`
	Time  time.Time `json:"time"`
	Title string    `json:"title"`
	//TODO add a Title field to Post
}

type Markdown string

func (m Markdown) MarshalJSON() ([]byte, error) {
	mkd := blackfriday.MarkdownCommon([]byte(m))

	js, err := json.Marshal(string(mkd))
	if err != nil {
		return nil, err
	}

	return js, nil
}

var db []Post

func main() {
	fmt.Println("Hello, World!")

	r := pat.New()

	r.Get("/hello", hello)
	r.Post("/markdown", markdown)
	r.Post("/posts", addPost)
	r.Get("/posts", getPosts)
	r.Delete("/posts/{id}", delPost)

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

func delPost(w http.ResponseWriter, r *http.Request) {
	//Figure out which post they want to delete
	idStr := r.URL.Query().Get(":id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id < 0 || id >= len(db) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db = append(db[:id], db[id+1:]...)

	w.WriteHeader(http.StatusNoContent)

}
