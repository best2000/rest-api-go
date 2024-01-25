package main

import (
	"database/sql"

	"github.com/best2000/rest-api-go/handler"
	"github.com/best2000/rest-api-go/repo"
	"github.com/best2000/rest-api-go/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewChiRouter(db *sql.DB) chi.Router {
	//init router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/"))
	r.Use(handler.AuthMid)

	dogRepo := repo.DogRepo{Db: db}
	dogHandler := handler.DogHandler{DogRepo: &dogRepo}
	r.Route("/dogs", func(r chi.Router) {
		r.Method("GET", "/{id}", util.HandlerFunc(dogHandler.GetDogByID))
		r.Method("GET", "/", util.HandlerFunc(dogHandler.ListAllDog))
		r.Method("POST", "/", util.HandlerFunc(dogHandler.CreateDog))
		r.Method("PATCH", "/{id}", util.HandlerFunc(dogHandler.UpdateDogByID))
		r.Method("DELETE", "/{id}", util.HandlerFunc(dogHandler.DeleteDogByID))
	})

	return r
}
