package api

import (
	"database/sql"

	"github.com/best2000/rest-api-go/api/handler"
	mid "github.com/best2000/rest-api-go/api/middleware"
	logging "github.com/best2000/rest-api-go/log"
	"github.com/best2000/rest-api-go/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewChiRouter(db *sql.DB, l *logging.Logger) chi.Router {
	//init router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/"))

	r.Use(mid.AuthMid)

	dogRepo := repo.DogRepo{Db: db, Log: l}
	dogHandler := handler.DogHandler{DogRepo: &dogRepo, Log: l}
	r.Route("/dogs", func(r chi.Router) {
		r.Get("/{id}", dogHandler.GetDogByID)
		r.Get("/", dogHandler.ListAllDog)
		r.Post("/", dogHandler.CreateDog)
		r.Patch("/{id}", dogHandler.UpdateDogByID)
		r.Delete("/{id}", dogHandler.DeleteDogByID)
	})

	

	return r
}