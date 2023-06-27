package router

import (
	"database/sql"
	"encoding/json"
	_ "go-setup/docs"
	"go-setup/internal/handler"
	"go-setup/internal/repository"
	"go-setup/internal/service"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{Message: "Hello World"})
	}).Methods("GET")
	// book api
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookServiceImpl(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	router.HandleFunc("/books/", bookHandler.FindAll).Methods("GET")
	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The path to your swagger.json file
	))
	return router
}
