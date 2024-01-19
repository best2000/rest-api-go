package main

import (
	"fmt"
	"net/http"

	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/db"
	"github.com/best2000/rest-api-go/handler"
	"github.com/best2000/rest-api-go/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config := config.GetConfig()

	db := database.NewPostgresDatabase(*config).Db
	defer db.Close()
	fmt.Println("connected to database.")

	//init router
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/"))

	dogRepo := repo.DogRepo{Db: db}
	dogHandler := handler.DogHandler{DogRepo: &dogRepo}
	r.Route("/dog", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/{id}", dogHandler.HandleGetDogByID)
		r.Get("/", dogHandler.HandleListAllDog)
		r.Post("/", dogHandler.HandleCreateDog)
		r.Patch("/{id}", dogHandler.HandleUpdateDogByID)
		r.Delete("/{id}", dogHandler.HandleDeleteDogByID)
	})

	fmt.Println("start server.")

	s := &http.Server{
		Addr:    config.App.Addr,
		Handler: r,
	}

	s.ListenAndServe()
}
