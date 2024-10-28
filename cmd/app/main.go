package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/File-Sharer/hash-generator-service/internal/config"
	"github.com/File-Sharer/hash-generator-service/internal/server"
	"github.com/File-Sharer/hash-generator-service/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := initEnv(); err != nil {
		logrus.Fatalf("error initializing env: %s", err.Error())
	}

	services := service.New()

	serverConfig := &config.GRPCServerConfig{
		Network: viper.GetString("app.network"),
		Addr: viper.GetString("app.addr"),
	}
	go func ()  {
		if err := server.RunGRPCServer(serverConfig, services); err != nil {
			logrus.Fatalf("error running grpc server: %s", err.Error())
		}
	}()

	logrus.Println("gRPC Server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("gRPC Server Shutting Down")
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initEnv() error {
	return godotenv.Load(".env")
}
