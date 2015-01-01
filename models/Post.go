package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	//	"time"
	"encoding/json"
)

var db *sql.DB

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}

type Poster interface {
	SavePost()
	DeletePost()
	UpdatePost(string, string)
}

func (post *Post) SavePost() {
	insquery := fmt.Sprintf(`insert into articles values ('{"id":"%s", "title":"%s", "content":"%s"}')`, post.Id, post.Title, post.Content)
	fmt.Printf(insquery)
	_, err := db.Query(insquery)
	if err != nil {
		log.Fatal(err)
	}
}

func (post *Post) DeletePost() {
	delquery := fmt.Sprintf(`delete from articles where data->>'id' = '%s'`, post.Id)
	fmt.Printf(delquery)
	_, err := db.Query(delquery)
	if err != nil {
		log.Fatal(err)
	}
}

func (post *Post) UpdatePost(title, content string) {
	updquery := fmt.Sprintf(`update articles set data->>'title' = '%s', data->>'content' = '%s' where data->>'id' = '%s'`, title, content, post.Id)
	fmt.Printf(updquery)
	_, err := db.Query(updquery)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDb(connStr string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		panic("Database error")
	}
	return db, err
}

func GetPosts(db *sql.DB, query string) map[string]*Post {
	var posts map[string]*Post
	posts = make(map[string]*Post, 0)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Can't do query --> %v\n", err)
		log.Fatal(err)
	}
	for rows.Next() {
		var data json.RawMessage
		if err := rows.Scan(&data); err != nil {
			log.Fatal(err)
		}
		post := new(Post)
		if err := json.Unmarshal(data, post); err != nil {
			log.Fatal(err)
		}
		posts[post.Id] = post
	}
	return posts
}

func InsertPost(db *sql.DB, post *Post) {
	id := post.Id
	title := post.Title
	content := post.Content
	insquery := fmt.Sprintf(`insert into articles values ('{"id":"%s", "title":"%s", "content":"%s"}')`, id, title, content)
	fmt.Printf(insquery)
	_, err := db.Query(insquery)
	if err != nil {
		log.Fatal(err)
	}
}

func DeletePost(db *sql.DB, post *Post) {
	delquery := fmt.Sprintf(`delete from articles where data->>'id' = '%s'`, post.Id)
	_, err := db.Query(delquery)
	if err != nil {
		log.Fatal(err)
	}
}
