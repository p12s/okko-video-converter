package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/p12s/okko-video-converter/api/pkg/broker"
	"github.com/p12s/okko-video-converter/api/pkg/handler"
	"github.com/p12s/okko-video-converter/api/pkg/repository"
	"github.com/p12s/okko-video-converter/api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("👺👺👺 error init config: %s\n", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("👺👺👺 error loading env variables: %s\n", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("👺👺👺 failed to initialize db: %s\n", err.Error())
	}

	// инит клин (репо-сервис-хендлер)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	kafka, err := broker.NewKafka()
	if err != nil {
		logrus.Fatalf("❌ kafka error: %s\n", err.Error())
	}
	go broker.Subscribe(kafka.Consumer, repos)

	handlers := handler.NewHandler(services, kafka)

	// ран сервер
	srv := new(Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while running http server: %s\n", err.Error())
		}
	}()
	logrus.Print("😀😀😀 service started 😀😀😀")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("🧟🧟🧟 Shutting Down 🧟🧟🧟")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

// initConfig - инициализация конфигов из configs/config
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

type Server struct {
	httpServer *http.Server
	kafka      *broker.Kafka
	doneC      chan struct{}
	closeC     chan struct{}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown - grace-full-выключение
func (s *Server) Shutdown(ctx context.Context) error {
	close(s.closeC)

	for {
		select {
		case <-ctx.Done():
			return errors.New("context.Done")
		case <-s.doneC:
			return nil
		}
	}
}
