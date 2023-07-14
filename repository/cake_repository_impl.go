package repository

import (
	"context"
	"database/sql"
	"errors"
	"felixsavero/cake-store-restful-api/helper"
	"felixsavero/cake-store-restful-api/model/domain"
)

type CakeRepositoryImpl struct {
}

func NewCakeRepository() CakeRepository {
	return &CakeRepositoryImpl{}
}

func (repository *CakeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake {
	query := "INSERT INTO cakes (title, description, rating, image) values (?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, cake.Title, cake.Description, cake.Rating, cake.Image)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cake.Id = int(id)
	return cake
}

func (repository *CakeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cake domain.Cake) domain.Cake {
	query := `UPDATE cakes SET title = ? ,
							description = ? ,
							rating = ? ,
							image = ? ,
							updated_at = current_timestamp
						WHERE id = ?`
	_, err := tx.ExecContext(ctx, query, cake.Title, cake.Description, cake.Rating, cake.Image, cake.Id)
	helper.PanicIfError(err)

	return cake
}

func (repository *CakeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, cake domain.Cake) {
	query := "UPDATE cakes SET deleted_at = current_timestamp WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, cake.Id)
	helper.PanicIfError(err)
}

func (repository *CakeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, cakeId int) (domain.Cake, error) {
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ? AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, query, cakeId)
	helper.PanicIfError(err)
	defer rows.Close()

	cake := domain.Cake{}
	if rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		helper.PanicIfError(err)
		return cake, nil
	} else {
		return cake, errors.New("cake not found")
	}
}

func (repository *CakeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Cake {
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE deleted_at IS NULL ORDER BY rating DESC, title ASC"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var cakes []domain.Cake
	for rows.Next() {
		cake := domain.Cake{}
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		helper.PanicIfError(err)
		cakes = append(cakes, cake)
	}
	return cakes
}
