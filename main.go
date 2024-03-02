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
	r := NewChiRouter(db)

	//server setup
	s := &http.Server{
		Addr:    config.App.Addr,
		Handler: r,
	}

	log.Info(`
 ______     ______     ______     ______      ______     ______   __      
/\  == \   /\  ___\   /\  ___\   /\__  _\    /\  __ \   /\  == \ /\ \     Nothing Special
\ \  __<   \ \  __\   \ \___  \  \/_/\ \/    \ \  __ \  \ \  _-/ \ \ \    Just a Prototype
 \ \_\ \_\  \ \_____\  \/\_____\    \ \_\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_/ /_/   \/_____/   \/_____/     \/_/      \/_/\/_/   \/_/     \/_/   EmptyMan_ 
      `)

	//start server
	err := s.ListenAndServe()
	
	panic(err.Error())
}
