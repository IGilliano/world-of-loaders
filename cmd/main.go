package main

import (
	"log"
	"net/http"
	"worldOfLoaders/pkg/handler"
	"worldOfLoaders/pkg/repository"
	"worldOfLoaders/pkg/service"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "456123789",
		DBName:   "wol",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Cant initialize db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	svc := service.NewService(rep)
	hdlr := handler.NewHandler(svc)
	if err = http.ListenAndServe("localhost:8080", hdlr.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err)
	}
}
