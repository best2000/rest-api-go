package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/best2000/rest-api-go/model"
	"github.com/best2000/rest-api-go/repo"
	"github.com/go-chi/chi/v5"
)

type DogHandler struct{
	DogRepo *repo.DogRepo
}

func (h *DogHandler) HandleCreateDog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleCreateDog")
	// dog := model.Dog{}
	// err := h.DogRepo.CreateDog(r.Context(), dog, nil)
}

func (h *DogHandler) HandleListAllDog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleListAllDog")
	w.Write([]byte("hello"))
}

func (h *DogHandler) HandleGetDogByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleGetDogByID")
	//parse id
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
	//get from db
	dog, err := h.DogRepo.GetDogById(r.Context(), id, nil)
	if err != nil {
        fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
	//encode json
	j, err := json.Marshal(dog)
	if err != nil {
        fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }

		w.Write(j)
}

func (h *DogHandler) HandleUpdateDogByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleUpdateDogByID")
	dog := model.Dog{}
	h.DogRepo.UpdateDogById(r.Context(), dog, nil)
}

func (h *DogHandler) HandleDeleteDogByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleDeleteDogByID")
	//parse id
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
	//del from db
	err = h.DogRepo.DeleteDogById(r.Context(), id, nil)
	if err != nil {
        fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
    }
}

func (h *DogHandler) HandleSwapDogNameByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleSwapDogNameByID")
	
	//begin txn
	txn, _ := h.DogRepo.Db.Begin()

	dog1 := model.Dog{}
	//update dog
	err := h.DogRepo.UpdateDogById(r.Context(), dog1, txn)
	if err != nil {
		txn.Rollback()
	}

	dog2 := model.Dog{}
	//update another dog
	err = h.DogRepo.UpdateDogById(r.Context(), dog2, txn)
	if err != nil {
		txn.Rollback()
	}
	
	err = txn.Commit()

	w.Write([]byte(err.Error()))
}