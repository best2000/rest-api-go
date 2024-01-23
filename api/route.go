package api

import (
	"database/sql"

	"github.com/best2000/rest-api-go/api/handler"
	mid "github.com/best2000/rest-api-go/api/middleware"
	"github.com/best2000/rest-api-go/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewChiRouter(db *sql.DB) chi.Router {
	//init router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/"))

	r.Use(mid.AuthMid)

	dogRepo := repo.DogRepo{Db: db}
	dogHandler := handler.DogHandler{DogRepo: &dogRepo}
	// r.With(mid.AuthMid).Route("/dogs", func(r chi.Router) {
		r.Get("/{id}", dogHandler.HandleGetDogByID)
		r.Get("/", dogHandler.HandleListAllDog)
		r.Post("/", dogHandler.HandleCreateDog)
		r.Patch("/{id}", dogHandler.HandleUpdateDogByID)
		r.Delete("/{id}", dogHandler.HandleDeleteDogByID)
	})

	

	return r
}