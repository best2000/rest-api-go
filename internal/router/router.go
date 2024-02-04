package router

import (
	"errors"
	"net/http"

	// "time"

	"rest-api/internal/handler"
	"rest-api/internal/buserror"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

func NewChiRouter() chi.Router {
	//init router
	r := chi.NewRouter()
	//middleware...
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

	r.Use(handler.PrePostMid)
	r.Use(handler.EndpointInfoMid)
	r.Use(handler.AuthMid)
	r.Use(handler.LogMid)

	// dogRepo := repo.DogRepo{Db: db}
	// dogHandler := handler.DogHandler{DogRepo: &dogRepo}

	r.Method("GET", "/ping", HandlerFunc(handler.Ping))

	return r
}

// custom handlerfunc that implement 'http.handler' interface
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//get logger from context
	log := r.Context().Value("logger").(*zap.Logger)

	//call main handler
	err := f(w, r)

	//error handling...
	if err != nil {
		log.Error(err.Error(), zap.Error(err))

		var busErr buserror.BusinessError
		//check bussiness error
		if errors.As(err, &busErr) {
			//set bussiness error response
			w.WriteHeader(busErr.HttpStatus)
			w.Write([]byte(busErr.Error()))
		} else {
			//std error
			//set 500 Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}
}
