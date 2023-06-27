package main

import (
	"go-setup/internal/config"
	"go-setup/internal/router"
	"go-setup/pkg/middleware"
	"go-setup/pkg/utils"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db := config.DatabaseConnection(cfg)
	defer db.Close()

	routers := router.SetupRouter(db)
	routers.Use(middleware.SetJSONContentType)
	log.Println("Server is running on port " + cfg.Server.Port)
	server := http.Server{Addr: cfg.Server.Host + ":" + cfg.Server.Port, Handler: routers}
	err := server.ListenAndServe()
	utils.PanicIfError(err)
}
