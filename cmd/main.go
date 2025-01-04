package main

import (
	todo "go-rest-exmpl"
	"go-rest-exmpl/pkg/handler"
	"go-rest-exmpl/pkg/repository"
	"go-rest-exmpl/pkg/service"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config initialization error %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Config initialization error %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})

	if err != nil {
		log.Fatalf("DB initialization error %s", err.Error())
	}

	repository := repository.NewRepositories(db)
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)
	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Occurred fatal error %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
