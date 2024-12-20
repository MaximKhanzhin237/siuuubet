package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Application struct {
}

var App = &Application{}

func main() {

	//создание сервера
	srv := &http.Server{
		Addr:    ":8080",
		Handler: App.routes(),
	}

	log.Println("Запуск сервера на 8080")
	err1 := srv.ListenAndServe()
	log.Fatal(err1)
}
