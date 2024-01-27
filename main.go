package main

import (
	"fmt"
	"net/http"

	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/db"
	"github.com/best2000/rest-api-go/logger"
)



func main() {
	//zap logger setup
	log := logger.New("dev")
	defer log.Sync()
	
	log.Info("initializing server...")

	//get config
	config := config.GetConfig()

	//connect db
	db := database.NewPostgresDatabase(*config).Db
	defer db.Close()
	log.Info("connected to database.")

	//TODO
	//pagination
	//add middleware pre/post handle, logger, request id
	//error/routes/logs middleware management
	//graceful shutdown

	log.Info(fmt.Sprintf("listening on %s.",config.App.Addr))

	//router setup
	r := NewChiRouter(db)

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
