package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/best2000/rest-api-go/log"
	"github.com/best2000/rest-api-go/model"
	"github.com/best2000/rest-api-go/repo"
	"github.com/go-chi/chi/v5"
)

type DogHandler struct{
	DogRepo *repo.DogRepo
	Log *logging.Logger
}

func (h *DogHandler) CreateDog(w http.ResponseWriter, r *http.Request) {
	h.Log.Info.Println("DogHandler HandleCreateDog")
	
	var dog model.DogCreateReq
	err := json.NewDecoder(r.Body).Decode(&dog)
	if err != nil {
        log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
	err = h.DogRepo.CreateDog(r.Context(), dog, nil)
	if err != nil {
        log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
}

func (h *DogHandler) ListAllDog(w http.ResponseWriter, r *http.Request) {
	h.Log.Info.Println("DogHandler HandleListAllDog")
	h.DogRepo.GetAllDog(r.Context(), nil)
	w.Write([]byte("lol"))
}

func (h *DogHandler) GetDogByID(w http.ResponseWriter, r *http.Request) {
	h.Log.Info.Println("DogHandler HandleGetDogByID")
	//parse id
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
	//get from db
	dog, err := h.DogRepo.GetDogById(r.Context(), id, nil)
	if err != nil {
        log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
	//encode json
	j, err := json.Marshal(dog)
	if err != nil {
        log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }

	w.Write(j)
}

func (h *DogHandler) UpdateDogByID(w http.ResponseWriter, r *http.Request) {
	h.Log.Info.Println("DogHandler HandleUpdateDogByID")
	dog := model.Dog{}
	h.DogRepo.UpdateDogById(r.Context(), dog, nil)
}

func (h *DogHandler) DeleteDogByID(w http.ResponseWriter, r *http.Request) {
	h.Log.Info.Println("DogHandler HandleDeleteDogByID")
	//parse id
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
	//del from db
	err = h.DogRepo.DeleteDogById(r.Context(), id, nil)
	if err != nil {
        log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
}