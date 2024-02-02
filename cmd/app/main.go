package main

import (
	"fmt"
	"net/http"

	"rest-api/internal/config"
	database "rest-api/internal/db"
	"rest-api/internal/logger"
	"rest-api/internal/router"
)

//TODO
//pagination
//graceful shutdown

func main() {
	//get config
	config := config.Load()

	//zap logger setup
	log := logger.Logger
	defer log.Sync()

	log.Info("initializing server...")

	//connect db
	db, err := database.Connect(*config)
	if err != nil {
		log.Error(err.Error())
	}
	defer db.Close()

	log.Info(fmt.Sprintf("listening on %s.", config.Server.Addr))

	//router setup
	r := router.NewChiRouter()

	//server setup
	s := &http.Server{
		Addr:    config.Server.Addr,
		Handler: r,
	}

	//start server
	go s.ListenAndServe()

	log.Info(`
 ______     ______     ______     ______      ______     ______   __      
/\  == \   /\  ___\   /\  ___\   /\__  _\    /\  __ \   /\  == \ /\ \     Nothing Special
\ \  __<   \ \  __\   \ \___  \  \/_/\ \/    \ \  __ \  \ \  _-/ \ \ \    Just a Prototype
 \ \_\ \_\  \ \_____\  \/\_____\    \ \_\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_/ /_/   \/_____/   \/_____/     \/_/      \/_/\/_/   \/_/     \/_/   EmptyMan_ 
      `)

	c := make(chan int)
	<-c
}
