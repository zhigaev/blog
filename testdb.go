package main

import (
	"./models"
	"fmt"
	"log"
)

const connStr = "host=192.168.75.128 port=5432 user=postgres dbname=blog sslmode=disable"

var QUERY_STR = "SELECT * from articles;"

func main() {
	db, err := ConnectDb(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	posts := GetPosts(db, QUERY_STR)
	for _, post := range posts {
		fmt.Printf("%s %s\n", post.Title, post.Content)
	}
	post := models.NewPost("algjalgj", "tavag", "bxbxnxn")
	InsertPost(db, post)
}
