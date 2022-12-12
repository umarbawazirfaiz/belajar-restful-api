package repository

import (
	"belajar-restful-api/exception"
	"belajar-restful-api/helper"
	"belajar-restful-api/model/domain"
	"context"
	"database/sql"
)

type categoryRepository struct {
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{}
}

func (repo *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		rows.Scan(category.GetId(), category.GetName())
		categories = append(categories, category)
	}

	return categories
}

func (repo *categoryRepository) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO category(name) VALUES(?)"
	res, err := tx.ExecContext(ctx, query, *category.GetName())
	helper.PanicIfError(err)

	lastId, _ := res.LastInsertId()
	id := int(lastId)
	category.SetId(&id)

	return category
}

func (repo *categoryRepository) Update(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "UPDATE category SET name=? WHERE id=?"
	_, err := tx.ExecContext(ctx, query, *category.GetName(), *category.GetId())
	helper.PanicIfError(err)
}

func (repo *categoryRepository) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM sscategory WHERE id=?"
	_, err := tx.ExecContext(ctx, query, *category.GetId())
	helper.PanicIfError(err)
}

func (repo *categoryRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Category {
	category := domain.Category{}
	query := "SELECT id, name FROM category WHERE id=?"
	row := tx.QueryRowContext(ctx, query, id)
	err := row.Scan(category.GetId(), category.GetName())
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	return category
}
