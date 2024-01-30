package router

import (
	"database/sql"
	// "time"

	"github.com/best2000/rest-api-go/internal/handler"
	"github.com/best2000/rest-api-go/internal/repo"
	"github.com/best2000/rest-api-go/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewChiRouter(db *sql.DB) chi.Router {
	//init router
	r := chi.NewRouter()
	//middleware...
	// r.Use(middleware.Timeout(time.Second*3))	//set 60 sec request context timeout 
	r.Use(cors.Handler(cors.Options{ //CORS
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// r.Use(middleware.Logger)
	r.Use(handler.PrePost)
	r.Use(handler.AuthMid)
	r.Use(middleware.Heartbeat("/"))
	

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
