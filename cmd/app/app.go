package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"hw_3/internal/config"
	"hw_3/internal/repository/postgres"
	"hw_3/internal/service"
	"hw_3/internal/service/domain"
	"hw_3/internal/transport/http/routes"
	"net/http"
	"time"
)

type App struct {
	cfg    *config.Config
	server *http.Server

	UserService domain.UserService
	UserRepo    domain.UserRepository
}

func NewApp(cfg *config.Config) *App {

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	userRepo := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	address := fmt.Sprintf(":%d", cfg.Http.Port)

	a := &App{
		cfg:         cfg,
		UserService: userService,
		UserRepo:    userRepo,
		server: &http.Server{
			Addr:           address,
			Handler:        initApi(userService),
			WriteTimeout:   5 * time.Second,
			ReadTimeout:    5 * time.Second,
			MaxHeaderBytes: 1 << 20, // 1 MB
		},
	}

	return a
}

func (a *App) Run() {
	go func() {
		err := a.server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()
}

func initApi(us domain.UserService) http.Handler {
	router := gin.Default()

	api := router.Group("/api")
	routes.InitRootRoute(api)
	routes.InitUserRoutes(api, us)

	return router
}

func (a *App) Stop(ctx context.Context) {
	// Место для различных функций нужных перед завершением сервера
	// (возможный сбор логов, работа с бд и прочее)

	// Говорим серверу не принимать новые запросы
	fmt.Println(a.server.Shutdown(ctx))
}

// Later
//func (a *App) InitAuth() {
//
//}
