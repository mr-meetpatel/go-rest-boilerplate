package service

import (
	"context"
	"go-setup/api/response"
	"go-setup/internal/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

// FindAll implements BookService
func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	//var bookResp []response.BookResponse
	bookResp := make([]response.BookResponse, 0)
	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Title: value.Title}
		bookResp = append(bookResp, book)
	}
	return bookResp

}
