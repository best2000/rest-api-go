package handler

import (
	"fmt"
	"net/http"

	"github.com/best2000/rest-api-go/model"
	"github.com/best2000/rest-api-go/repo"
)

type DogHandler struct{
	dogRepo *repo.DogRepo
}

func (h *DogHandler) HandleCreateDog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleCreateDog")
}

func (h *DogHandler) HandleListAllDog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleListAllDog")
}

func (h *DogHandler) HandleGetDogByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleGetDogByID")
}

func (h *DogHandler) HandleUpdateDogByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleUpdateDogByID")
}

func (h *DogHandler) HandleDeleteDogByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleDeleteDogByID")
}

func (h *DogHandler) HandleSwapDogNameByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DogHandler HandleSwapDogNameByID")
	//begin txn
	txn, _ := h.dogRepo.Db.BeginTx(r.Context(), nil)

	txn.Exec("SELECT pg_sleep(10)")

	//update dog
	h.dogRepo.UpdateDogById(nil, model.Dog{}, txn)

	//update another dog
	h.dogRepo.UpdateDogById(nil, model.Dog{}, txn)
	
	//commit txn
	txn.Commit()
}