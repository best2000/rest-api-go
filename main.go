package main

import (
	"fmt"
	"net/http"

	application "github.com/best2000/rest-api-go/app"
	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/db"
	"github.com/best2000/rest-api-go/handler"
	"github.com/best2000/rest-api-go/repo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-delve/delve/service/api"
	"golang.org/x/tools/cmd/getgo/server"
)

func main() {
	config := config.GetConfig()

	db := database.NewPostgresDatabase(*config).Db
	defer db.Close()
	fmt.Println("connected to database.")


	//pagination
	//add middleware pre/post handle, logger, request id 
	//error/routes/logs middleware management
	//graceful shutdown

	

	fmt.Println("start server "+config.App.Addr)
	
	//load route
	r := api.NewChiRouter(db)

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
