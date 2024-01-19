package application

import (
	"github.com/best2000/golang-chi-api/handler"
	"github.com/best2000/golang-chi-api/repository/order"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()
	
	router.Use(middleware.Heartbeat("/"))
	router.Use(middleware.Logger)

	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
  router.Delete("/{id}", orderHandler.DeleteByID)
}
