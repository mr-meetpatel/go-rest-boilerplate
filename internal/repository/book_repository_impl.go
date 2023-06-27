package repository

import (
	"context"
	"database/sql"
	"go-setup/internal/model"
	"go-setup/pkg/utils"
)

type BookRepositoryImpl struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{db: db}
}

// FindAll implements BookRepository
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	tx, err := b.db.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	SQL := "SELECT id,title,author FROM books"
	result, errQuery := tx.QueryContext(ctx, SQL)
	utils.PanicIfError(errQuery)
	defer result.Close()

	books := make([]model.Book, 0)
	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Title)
		utils.PanicIfError(err)

		books = append(books, book)
	}
	return books
}
