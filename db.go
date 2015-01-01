package main

import (
	"./models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	//	"time"
	"encoding/json"
)

func ConnectDb(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		panic("Database error")
	}
	return db, err
}

func GetPosts(db *sql.DB, query string) map[string]*models.Post {
	var posts map[string]*models.Post
	posts = make(map[string]*models.Post, 0)
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
		post := new(models.Post)
		if err := json.Unmarshal(data, post); err != nil {
			log.Fatal(err)
		}
		posts[post.Id] = post
	}
	return posts
}

func InsertPost(db *sql.DB, post *models.Post) {
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

func DeletePost(db *sql.DB, post *models.Post) {
	delquery := fmt.Sprintf(`delete from articles where data->>'id' = '%s'`, post.Id)
	_, err := db.Query(delquery)
	if err != nil {
		log.Fatal(err)
	}
}
