package main

import (
	"fmt"
	"net/http"

	"github.com/best2000/rest-api-go/api"
	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/db"
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

	fmt.Println("start server " + config.App.Addr)

	//load route
	r := api.NewChiRouter(db)

	s := &http.Server{
		Addr:    config.App.Addr,
		Handler: r,
	}

	go s.ListenAndServe()

	fmt.Println(`
 ______     ______     ______     ______      ______     ______   __      
/\  == \   /\  ___\   /\  ___\   /\__  _\    /\  __ \   /\  == \ /\ \     Nothing Special
\ \  __<   \ \  __\   \ \___  \  \/_/\ \/    \ \  __ \  \ \  _-/ \ \ \    It's Just a Prototype
 \ \_\ \_\  \ \_____\  \/\_____\    \ \_\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_/ /_/   \/_____/   \/_____/     \/_/      \/_/\/_/   \/_/     \/_/   EmptyMan_ 
      `)

	c := make(chan int)
	<-c
}
