package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, _ = sql.Open("mysql", "go_course:go_course@/go_course")

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{id}/view", ViewHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))
	if err := t.ExecuteTemplate(w, "layout.html", ListPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/view.html"))
	err := t.ExecuteTemplate(w, "layout.html", GetPostById(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetPostById(id int) Post {
	row := db.QueryRow("select * from posts where id = ?", id)
	post := Post{}
	err := row.Scan(&post.Id, &post.Title, &post.Body)
	if err != nil {
		return Post{}
	}
	return post
}

func ListPosts() []Post {
	rows, err := db.Query("select * from posts")
	checkErr(err)
	var items []Post
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Body)
		if err != nil {
			return nil
		}
		items = append(items, post)
	}
	return items
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
