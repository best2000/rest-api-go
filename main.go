package main

import (
	"fmt"
	"net/http"

	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/db"
	"github.com/best2000/rest-api-go/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config := config.GetConfig()

	db := database.NewPostgresDatabase(*config)
	defer db.Close()
	fmt.Println("connected to " + db.GetDbSysInfo() + ".")

	//init router
	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	dataHandler := handler.DataHandler{} 
	r.Route("/data", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/", dataHandler.List)
		r.Post("/", dataHandler.Create)
	})

	s := &http.Server{
		Addr:    config.App.Addr,
		Handler: r,
	}

	s.ListenAndServe()
}
