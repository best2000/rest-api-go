package handler

import (
	"fmt"
	"net/http"

	"github.com/best2000/rest-api-go/model"
	"github.com/best2000/rest-api-go/repo"
)

type DogHandler struct{
	DogRepo *repo.DogRepo
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
	txn, _ := h.DogRepo.Db.BeginTx(r.Context(), nil)

	txn.Exec("SELECT pg_sleep(10)")

	dog1 := model.Dog{}
	//update dog
	h.DogRepo.UpdateDogById(r.Context(), dog1, txn)

	dog2 := model.Dog{}
	//update another dog
	h.DogRepo.UpdateDogById(r.Context(), dog2, txn)
	
	//commit txn
	txn.Commit()
	
	w.Write([]byte("done"))
}