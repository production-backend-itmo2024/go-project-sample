package main

import (
	"github.com/joho/godotenv"
	"go-sample/internal/domain"
	"go-sample/internal/pkg/db_postgres"
	"go-sample/internal/repository"
	"log"
	"os"

	"github.com/spf13/viper"
	"go-sample/internal/app"
	"go-sample/internal/pkg/client/weather"
	"go-sample/internal/pkg/utils/server"

	_ "github.com/lib/pq"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	cfg := db_postgres.Config{
		Password: os.Getenv("DB_PASSWORD"),
	}
	cfg.Host = viper.GetString("db.host")
	cfg.Port = viper.GetString("db.port")
	cfg.User = viper.GetString("db.user")
	cfg.Dbname = viper.GetString("db.dbname")
	cfg.SSL = viper.GetString("db.ssl")

	db, err := db_postgres.NewPostrgresDB(cfg)

	if err != nil {
		log.Fatal(err.Error())
	}

	// DAL
	repo := repository.NewRepository(db)
	client := weather.NewClient()
	// Domain
	serv := domain.NewService(client, repo)
	// Handlers
	handler := app.NewHandler(serv)

	servr := server.NewServer(viper.GetString("port"), handler.InitRoutes())

	if err := servr.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
