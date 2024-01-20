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


	//add middlewear pre/post handle, logger, request id 
	//init router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/"))
	

	dogRepo := repo.DogRepo{Db: db}
	dogHandler := handler.DogHandler{DogRepo: &dogRepo}
	r.Route("/dogs", func(r chi.Router) {
		r.Get("/{id}", dogHandler.HandleGetDogByID)
		r.Get("/", dogHandler.HandleListAllDog)
		r.Post("/", dogHandler.HandleCreateDog)
		r.Patch("/{id}", dogHandler.HandleUpdateDogByID)
		r.Delete("/{id}", dogHandler.HandleDeleteDogByID)
	})

	fmt.Println("start server "+config.App.Addr)

	s := &http.Server{
		Addr:    config.App.Addr,
		Handler: r,
	}

	go s.ListenAndServe()

	fmt.Println(`
 ______     ______     ______     ______      ______     ______   __    
/\  == \   /\  ___\   /\  ___\   /\__  _\    /\  __ \   /\  == \ /\ \   
\ \  __<   \ \  __\   \ \___  \  \/_/\ \/    \ \  __ \  \ \  _-/ \ \ \  
 \ \_\ \_\  \ \_____\  \/\_____\    \ \_\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_/ /_/   \/_____/   \/_____/     \/_/      \/_/\/_/   \/_/     \/_/`)
	
	c := make(chan int)
	<-c
}
