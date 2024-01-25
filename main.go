package main

import (
	"log"
	"net/http"

	"github.com/best2000/rest-api-go/api"
	"github.com/best2000/rest-api-go/config"
	"github.com/best2000/rest-api-go/db"
	logging "github.com/best2000/rest-api-go/log"
)

func main() {
	//logger setup
	log := logging.NewLogger(log.Ldate | log.Ltime | log.Lshortfile)
	

	config := config.GetConfig()

	db := database.NewPostgresDatabase(*config).Db
	defer db.Close()
	log.Info.Println("connected to database.")

	//pagination
	//add middleware pre/post handle, logger, request id
	//error/routes/logs middleware management
	//graceful shutdown

	log.Info.Printf("start server at %s.\n", config.App.Addr)

	//load route
	r := api.NewChiRouter(db, log)

	s := &http.Server{
		Addr:    config.App.Addr,
		Handler: r,
	}

	go s.ListenAndServe()

	log.Info.Println(`
 ______     ______     ______     ______      ______     ______   __      
/\  == \   /\  ___\   /\  ___\   /\__  _\    /\  __ \   /\  == \ /\ \     Nothing Special
\ \  __<   \ \  __\   \ \___  \  \/_/\ \/    \ \  __ \  \ \  _-/ \ \ \    Just a Prototype
 \ \_\ \_\  \ \_____\  \/\_____\    \ \_\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_/ /_/   \/_____/   \/_____/     \/_/      \/_/\/_/   \/_/     \/_/   EmptyMan_ 
      `)

	c := make(chan int)
	<-c
}
