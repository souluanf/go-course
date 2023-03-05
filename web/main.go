package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, _ = sql.Open("mysql", "go_course:go_course@/go_course")

func main() {

	stmt, err := db.Prepare("insert into posts(title, body) values(?, ?)")
	checkErr(err)

	_, err = stmt.Exec("My first post", "This is my first post")
	checkErr(err)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		post := Post{Id: 0, Title: "Unamed Post", Body: "Our contents"}

		if title := request.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(writer, "index.html", post); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
