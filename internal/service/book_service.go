package service

import (
	"context"
	"go-setup/api/response"
)

type BookService interface {
	FindAll(ctx context.Context) []response.BookResponse
}
