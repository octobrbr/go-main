package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/mongo"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.

	var posts = []storage.Post{
		{
			ID:          1,
			Title:       "How to Build Your First Web Application with Go",
			Content:     "In this article, you will discover how to leverage Go for web development by building a news application in the language.",
			AuthorID:    0,
			AuthorName:  "Ayooluwa Isaiah",
			CreatedAt:   time.Now().Unix(),
			PublishedAt: 0,
		},
		{
			ID:          2,
			Title:       "How to Use MongoDB with Go",
			Content:     "This tutorial will teach you how to use MongoDB databases with the Go programming language by connecting to your MongoDB Atlas cluster.",
			AuthorID:    1,
			AuthorName:  "Ukeje Goodness",
			CreatedAt:   time.Now().Unix(),
			PublishedAt: 0,
		},
		{
			ID:          3,
			Title:       "Using MongoDB with Docker",
			Content:     "In this article, you’ll learn the best practices for running a MongoDB container. You’ll also learn how to host a simple Flask app and how to use Docker volumes to persist data in a Docker container.",
			AuthorID:    2,
			AuthorName:  "Soumi Bardhan",
			CreatedAt:   time.Now().Unix(),
			PublishedAt: 0,
		},
		{
			ID:          4,
			Title:       "Building a simple app with Go and PostgreSQL",
			Content:     "In this article, we will be looking at how we can use Postgres in a Go application.",
			AuthorID:    3,
			AuthorName:  "Emmanuel John",
			CreatedAt:   time.Now().Unix(),
			PublishedAt: 0,
		},
	}

	// Создаём объекты баз данных.

	// БД в памяти.
	//db := memdb.New()

	// Реляционная БД PostgreSQL.
	/*pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}

	constr := "postgres://brannon:" + pwd + "@localhost/news"
	db, err := postgres.New(constr)
	if err != nil {
		log.Fatal(err)
	}*/

	// Документная БД MongoDB.
	db, err := mongo.New("mongodb://127.0.0.1:27017/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()

	for _, post := range posts {
		err = db.AddPost(post)
		if err != nil {
			fmt.Printf("Error %d", err)
			return
		}
	}
	fmt.Println("      Added posts")
	posts, err = db.Posts()
	if err != nil {
		fmt.Printf("Error %d", err)
		return
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	for _, post := range posts {
		post.Title = strconv.Itoa(post.ID) + ". " + post.Title
		err = db.UpdatePost(post)
		if err != nil {
			fmt.Printf("Error %d", err)
		}
	}
	fmt.Println("      Updated posts")
	posts, err = db.Posts()
	if err != nil {
		fmt.Printf("Error %d", err)
		return
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	for _, post := range posts {
		if post.ID%2 > 0 {
			err = db.DeletePost(post)
			if err != nil {
				fmt.Printf("Error %d", err)
			}
		}

	}
	fmt.Println("      After delete posts")
	posts, err = db.Posts()
	if err != nil {
		fmt.Printf("Error %d", err)
		return
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	/*
		var srv server

		// Инициализируем хранилище сервера конкретной БД.
		srv.db = db3

		// Создаём объект API и регистрируем обработчики.
		srv.api = api.New(srv.db)

		// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
		// Предаём серверу маршрутизатор запросов,
		// поэтому сервер будет все запросы отправлять на маршрутизатор.
		// Маршрутизатор будет выбирать нужный обработчик.
		http.ListenAndServe(":8080", srv.api.Router())
	*/
}
