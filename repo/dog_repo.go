package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/best2000/rest-api-go/logger"
	"github.com/best2000/rest-api-go/model"
	"github.com/best2000/rest-api-go/util"
)

type DogRepo struct {
	Db *sql.DB
}

func (r *DogRepo) CreateDog(ctx context.Context, d model.DogCreateReq, db util.DbExecutor) error {

	if db == nil {
		db = r.Db
	}
	sql := fmt.Sprintf("INSERT INTO dogs (name, breed) VALUES ('%s', '%s');", d.Name, d.Breed)
	_, err := db.ExecContext(ctx, sql)
	return err
}

func (r *DogRepo) GetAllDog(ctx context.Context, db util.DbQuerist) error {
	log := logger.FromCtx(ctx)
	if db == nil {
		db = r.Db
	}

	sql := "SELECT pg_sleep(1);"
	log.Info("sql: " + sql)
	
	log.Info("Just sleep...")

	_, err := r.Db.ExecContext(ctx, sql)
	return err
}

func (r *DogRepo) GetDogById(ctx context.Context, id int, db util.DbQuerist) (model.Dog, error) {
	if db == nil {
		db = r.Db
	}

	sql := fmt.Sprintf("SELECT id, name, breed, created_at, updated_at FROM dogs WHERE id = '%d';", id)
	row := db.QueryRowContext(ctx, sql)

	dog := model.Dog{}
	err := row.Scan(&dog.Id, &dog.Name, &dog.Breed, &dog.CreatedAt, &dog.UpdatedAt)
	return dog, err
}

func (r *DogRepo) UpdateDogById(ctx context.Context, d model.Dog, db util.DbExecutor) error {
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

func (r *DogRepo) DeleteDogById(ctx context.Context, id int, db util.DbExecutor) error {
	if db == nil {
		db = r.Db
	}
	sql := fmt.Sprintf("DELETE FROM dogs WHERE id = '%d'", id)
	_, err := db.ExecContext(ctx, sql)
	return err
}
