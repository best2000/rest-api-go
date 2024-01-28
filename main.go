package main

import (
	"fmt"
	"net/http"

	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/database"
	"github.com/best2000/rest-api-go/logger"
)

//TODO
//pagination
//graceful shutdown
//path check in middleware to auth or log/audit/pii
	
func main() {
	//zap logger setup
	log := logger.New("dev")
	defer log.Sync()
	
	log.Info("initializing server...")

	//get config
	config := config.Load()

	//connect db
	db := database.New(*config).Db
	defer db.Close()
	log.Info("connected to database.")

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
