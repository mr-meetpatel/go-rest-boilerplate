package handler

import (
	_ "go-setup/api/response"
	"go-setup/internal/service"
	"go-setup/pkg/utils"
	"net/http"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{BookService: bookService}
}

// GetOrders godoc
// @Summary Get details of all books
// @Description Get details of all books
// @Tags Book
// @Accept  json
// @Produce  json
// @Success 200 {array} response.BookResponse
// @Router /books [get]
func (handler *BookHandler) FindAll(writer http.ResponseWriter, requests *http.Request) {
	result := handler.BookService.FindAll(requests.Context())
	utils.WriteResponseBody(writer, result)
}
