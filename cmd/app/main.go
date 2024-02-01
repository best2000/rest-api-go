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
	log := logger.New(config.App.Env)
	defer log.Sync()

	log.Info("initializing server...")

	//connect db
	db := database.New(*config).Db
	defer db.Close()
	log.Info("connected to database.")

	log.Info(fmt.Sprintf("listening on %s.", config.App.Addr))

	//router setup
	r := router.NewChiRouter(db)

	//server setup
	s := &http.Server{
		Addr:    config.App.Addr,
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
