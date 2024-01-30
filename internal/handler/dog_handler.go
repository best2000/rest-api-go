package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/best2000/rest-api-go/logger"
	"github.com/best2000/rest-api-go/model"
	"github.com/best2000/rest-api-go/repo"
	"github.com/best2000/rest-api-go/util"
	"github.com/go-chi/chi/v5"
)

type DogHandler struct {
	DogRepo *repo.DogRepo
}

func (h *DogHandler) CreateDog(w util.ResponseWriter, r *http.Request) error {
	var dog model.DogCreateReq
	err := json.NewDecoder(r.Body).Decode(&dog)
	if err != nil {
		return err
	}

	err = h.DogRepo.CreateDog(r.Context(), dog, nil)
	if err != nil {
		return err
	}
	return nil
}

func (h *DogHandler) ListAllDog(w util.ResponseWriter, r *http.Request) error {
	log := logger.FromCtx(r.Context())

	log.Info("failing...")
	
	return h.DogRepo.GetAllDog(r.Context(), nil)
}

func (h *DogHandler) GetDogByID(w util.ResponseWriter, r *http.Request) error {
	//parse id
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}
	//get from db
	dog, err := h.DogRepo.GetDogById(r.Context(), id, nil)
	if err != nil {
		return err
	}

	return w.WriteAsJson(dog)
}

func (h *DogHandler) UpdateDogByID(w util.ResponseWriter, r *http.Request) error {
	dog := model.Dog{}
	h.DogRepo.UpdateDogById(r.Context(), dog, nil)
	return nil
}

func (h *DogHandler) DeleteDogByID(w util.ResponseWriter, r *http.Request) error {
	//parse id
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}
	//del from db
	err = h.DogRepo.DeleteDogById(r.Context(), id, nil)
	if err != nil {
		return err
	}
	return nil
}
