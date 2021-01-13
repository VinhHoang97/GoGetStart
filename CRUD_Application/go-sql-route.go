package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "testdb"
	dbPass = "Aa123456"
	dbHost = "localhost"
	dbPort = "3306"
)

func routers() *chi.Mux {
	router.Get("/posts", AllPosts)
	router.Get("/posts/{id}", DetailPost)
	router.Post("/posts", CreatePost)
	router.Put("/posts/{id}", UpdatePost)
	router.Delete("/posts/{id}", DeletePost)

	return router
}
