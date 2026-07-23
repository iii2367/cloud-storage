package main

import (
	authRepoPostgres "cloud-storage/internal/auth/repository/postgres"
	authService "cloud-storage/internal/auth/service"
	authHttp "cloud-storage/internal/auth/transport/http"

	storageRepoPostgres "cloud-storage/internal/storage/repository/postgres"
	storageService "cloud-storage/internal/storage/service"
	storageHttp "cloud-storage/internal/storage/transport/http"

	"cloud-storage/internal/config"
	"cloud-storage/internal/database"
	"cloud-storage/internal/jwt"
	"cloud-storage/internal/web"

	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad("MAIN")

	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	jwtManager := jwt.New(cfg.JWT)

	userRepo := authRepoPostgres.NewUserRepository(db)
	sessionRepo := authRepoPostgres.NewSessionRepository(db)
	authService := authService.NewService(userRepo, sessionRepo, jwtManager)
	authHandler := authHttp.NewHandler(authService)
	authMiddleware := authHttp.NewMiddleware(jwtManager)

	treeNodeRepo := storageRepoPostgres.NewTreeNodeRepository(db)
	storageService := storageService.NewService(treeNodeRepo, jwtManager)
	storageHandler := storageHttp.NewHandler(storageService)
	storageMiddleware := storageHttp.NewMiddleware(jwtManager)

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*")
	router.Static("/css", "./web/static/css")
	router.Static("/js", "./web/static/js")

	authHttp.RegisterRoutes(router, authHandler, authMiddleware)
	storageHttp.RegisterRoutes(router, storageHandler, storageMiddleware)

	webHandler := web.NewHandler()
	web.RegisterRoutes(router, webHandler)

	addr := fmt.Sprintf("%s:%s",
		cfg.Server.Host,
		cfg.Server.Port,
	)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		log.Println("server started on", addr)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf(
			"server forced shutdown: %v",
			err,
		)
	}

	log.Println("server stopped gracefully")
}
