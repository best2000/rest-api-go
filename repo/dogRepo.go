package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/best2000/rest-api-go/model"
	"github.com/best2000/rest-api-go/util"
)

type DogRepo struct {
	Db *sql.DB
}

func (r *DogRepo) CreateDog(ctx context.Context, d model.Dog, db util.DbExecutor) error {
	fmt.Println("DogRepo CreateDog")
	if db == nil {
		db = r.Db	
	}
	sql := fmt.Sprintf("INSERT INTO dogs(name, breed) VALUES (%s, %s);", d.Name, d.Breed)
	_, err := db.ExecContext(ctx, sql)
	return err
}

func (r *DogRepo) GetAllDog(ctx context.Context, exe util.DbExecutor) error {
	fmt.Println("DogRepo GetAllDog")
	return nil
}

func (r *DogRepo) GetDogById(ctx context.Context, id uint32, db util.DbQuerist) (model.Dog, error) {
	if db == nil {
		db = r.Db	
	}
	fmt.Println("DogRepo GetDogById")
	sql := fmt.Sprintf("SELECT * FROM dogs WHERE id = %d;", id)
	_, err := db.QueryContext(ctx, sql)
	dog := model.Dog{}
	return dog, err
}

func (r *DogRepo) UpdateDogById(ctx context.Context , d model.Dog, db util.DbExecutor) error {
	fmt.Println("DogRepo UpdateDogById")
	if db == nil {
		db = r.Db	
	}
	sql := fmt.Sprintf(`
		UPDATE dogs 
		SET name  = %s
			breed = %s	
		WHERE id = %d;`,
		d.Name, d.Breed, d.Id,
	)
	_, err := db.ExecContext(ctx, sql)
	return err
}

func (r *DogRepo) DeleteDogById(ctx context.Context, id uint32, db util.DbExecutor) error {
	fmt.Println("DogRepo DeleteDogById")
	if db == nil {
		db = r.Db	
	}
	sql := fmt.Sprintf("DELETE FROM dogs WHERE id = %d", id)
	_, err := db.ExecContext(ctx, sql)
	return err
}