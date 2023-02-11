package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rovoq19/GoChat/cmd/ws"
	"github.com/rovoq19/GoChat/pkg/handler"
	"github.com/rovoq19/GoChat/pkg/repository"
	"github.com/rovoq19/GoChat/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type wsRunner struct {
	upgrader *websocket.Upgrader
	wsServer *http.Server
	hub      *ws.Hub
}

func newWsRunner() *wsRunner {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	hub := ws.NewHub()

	return &wsRunner{upgrader: upgrader, hub: hub}
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	wsRunner := newWsRunner()
	repos := repository.NewRepository(db)
	services := service.NewService(repos, wsRunner.hub)
	handlers := handler.NewHandler(services)

	srv := new(Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Printf("HTTP server started on port %s", viper.GetString("port"))

	wsHandler := wsRunner.getWSHandler()
	wsRunner.wsServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("wsPort")),
		Handler: wsHandler,
	}
	go func() {
		if err := wsRunner.wsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running websocket server: %s", err.Error())
		}
	}()
	logrus.Printf("Websocket server started on port %s", viper.GetString("wsPort"))

	wsRunner.hub.Run()

	logrus.Print("GoChat Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("GoChat Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (run *wsRunner) getWSHandler() http.Handler {
	ginHandler := gin.New()

	ginHandler.GET("/", func(c *gin.Context) {
		run.serveWS(c.Writer, c.Request)
	})

	return ginHandler
}

func (run *wsRunner) serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := run.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := ws.NewClient(run.hub, conn)
	client.Serve()
}
