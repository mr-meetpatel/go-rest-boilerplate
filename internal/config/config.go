package config

import (
	"database/sql"
	"fmt"
	"go-setup/pkg/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg := &Config{}

	cfg.Server.Host = os.Getenv("SERVER_HOST")
	cfg.Server.Port = os.Getenv("SERVER_PORT")

	cfg.Database.Host = os.Getenv("DB_HOST")
	cfg.Database.Port = os.Getenv("DB_PORT")
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.Name = os.Getenv("DB_NAME")

	validateConfig(cfg)

	return cfg
}

func validateConfig(cfg *Config) {
	// Add your validation logic here
	if cfg.Server.Host == "" {
		log.Fatal("Missing SERVER_HOST environment variable")
	}

	if cfg.Server.Port == "" {
		log.Fatal("Missing SERVER_PORT environment variable")
	}

	if cfg.Database.Host == "" {
		log.Fatal("Missing DB_HOST environment variable")
	}

	if cfg.Database.Port == "" {
		log.Fatal("Missing DB_PORT environment variable")
	}

	if cfg.Database.User == "" {
		log.Fatal("Missing DB_USER environment variable")
	}

	if cfg.Database.Password == "" {
		log.Fatal("Missing DB_PASSWORD environment variable")
	}

	if cfg.Database.Name == "" {
		log.Fatal("Missing DB_NAME environment variable")
	}
}

func DatabaseConnection(cfg *Config) *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", dbInfo)
	utils.PanicIfError(err)
	err = db.Ping()
	utils.PanicIfError(err)
	return db
}
