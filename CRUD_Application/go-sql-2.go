package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8",  dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	catch(err)
}