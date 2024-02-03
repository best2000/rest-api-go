package main

import (
	"fmt"
	"net/http"

	"rest-api/internal/config"
	"rest-api/internal/db"
	"rest-api/internal/logger"
	"rest-api/internal/router"
	"rest-api/internal/token"
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

	//setup jwt auth
	tkn.New(*config)

	//connect db
	db, err := db.Connect(*config)
	if err != nil {
		log.Error(err.Error())
	}
	defer db.Close()

	//router setup
	r := router.NewChiRouter()

	//server setup
	s := &http.Server{
		Addr:    config.Server.Addr,
		Handler: r,
	}

	log.Info(fmt.Sprintf("listening on %s.", config.Server.Addr))

	//start server
	s.ListenAndServe()
}
