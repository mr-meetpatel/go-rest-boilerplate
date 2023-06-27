package repository

import (
	"context"
	"go-setup/internal/model"
)

type BookRepository interface {
	FindAll(ctx context.Context) []model.Book
}
