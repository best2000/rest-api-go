package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/db"
)

func main() {
	//logger setup
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)	
	logger := slog.Default()
	slog.SetDefault(logger)
	


	config := config.GetConfig()

	db := database.NewPostgresDatabase(*config).Db
	defer db.Close()
	slog.Info("connected to database.")

	//pagination
	//add middleware pre/post handle, logger, request id
	//error/routes/logs middleware management
	//graceful shutdown

	slog.Info("start server at "+ config.App.Addr + ".")

	//load route
	r := NewChiRouter(db)

	s := &http.Server{
		Addr:    config.App.Addr,
		Handler: r,
	}

	go s.ListenAndServe()

	slog.Info(`
 ______     ______     ______     ______      ______     ______   __      
/\  == \   /\  ___\   /\  ___\   /\__  _\    /\  __ \   /\  == \ /\ \     Nothing Special
\ \  __<   \ \  __\   \ \___  \  \/_/\ \/    \ \  __ \  \ \  _-/ \ \ \    Just a Prototype
 \ \_\ \_\  \ \_____\  \/\_____\    \ \_\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_/ /_/   \/_____/   \/_____/     \/_/      \/_/\/_/   \/_/     \/_/   EmptyMan_ 
      `)

	c := make(chan int)
	<-c
}
