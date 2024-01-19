package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/best2000/rest-api-go/model"
)

type DogRepo struct {
	Db *sql.DB
}

func (r *DogRepo) CreateDog(ctx context.Context, d model.Dog, txn *sql.Tx) error {
	fmt.Println("DogRepo CreateDog")
	return nil
}

func (r *DogRepo) GetAllDog(ctx context.Context, txn *sql.Tx) error {
	fmt.Println("DogRepo GetAllDog")
	return nil
}

func (r *DogRepo) GetDogById(ctx context.Context, id uint32, txn *sql.Tx) error {
	fmt.Println("DogRepo GetDogById")
	return nil
}

func (r *DogRepo) UpdateDogById(ctx context.Context, d model.Dog, txn *sql.Tx) error {
	fmt.Println("DogRepo UpdateDogById")
	if txn != nil {
		c := txn
		c.Exec("SELECT pg_sleep(10)")
	} else {
		c := r.Db
		c.ExecContext(ctx,"SELECT pg_sleep(10)")
	}

	return nil
}

func (r *DogRepo) DeleteDogById(ctx context.Context, id uint32, txn *sql.Tx) error {
	fmt.Println("DogRepo DeleteDogById")
	return nil
}