package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mihail-1212/todo-project-backend/internal/config"
	"github.com/mihail-1212/todo-project-backend/internal/delivery/http"
	"github.com/mihail-1212/todo-project-backend/internal/repository"
	"github.com/mihail-1212/todo-project-backend/internal/repository/postgres"
	"github.com/mihail-1212/todo-project-backend/internal/service"
	"github.com/mihail-1212/todo-project-backend/pkg/auth"
	authRepository "github.com/mihail-1212/todo-project-backend/pkg/auth/repository"
	"github.com/mihail-1212/todo-project-backend/pkg/logger"
	serverPackage "github.com/mihail-1212/todo-project-backend/pkg/server"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title TODO list project API
// @version 1.0
// @description TODO API for SPA project

// @contact.name API Support
// @contact.email m.a.mokruschin@yandex.com

// @license.name MIT

// @host localhost:8090
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Инициализация логгера
	log, _ := logger.NewLoggerDev()

	// Инициализация конфига
	if err := config.InitConfig(); err != nil {
		log.Panic("error initializing configs",
			zap.String("package", "main"),
			zap.String(" function", "initConfig"),
			zap.Error(err))
	}

	// Подключение к postgres
	db, err := postgres.NewPostgresDB(postgres.DBConfig{
		Host:     viper.GetString("db.host"),
		Password: viper.GetString("db.password"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Panic("error initializing database",
			zap.String("package", "main"),
			zap.String("function", "postgres.NewPostgresDB"),
			zap.Error(err))
	}

	// Repository
	repo := repository.NewRepository(db)
	// Authorizer
	authorizerRepo := authRepository.NewRepository(repo.User)
	authorizer := auth.NewAuthorizer(
		authorizerRepo,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		time.Hour,
	)
	// Services
	service := service.NewServices(repo)
	// Handler
	handler := http.NewHandler(service, authorizer)
	server := new(serverPackage.Server)

	port := viper.GetString("port")
	router := handler.InitAPI()

	// Запуск сервера
	go func() {
		serverCfg := serverPackage.NewServerConfig(port, router)
		err = server.Run(serverCfg)
		if err != nil {
			log.Info("error occured while running http server",
				zap.String("package", "main"),
				zap.String("function", "server.Run"),
				zap.Error(err))
		}
	}()
	log.Info("Start listening server...")
	log.Info("http://localhost:" + port + "/")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Server shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Error("error occured on server shutting down", zap.Error(err))
	}
}
