package main

import (
	"fmt"
	"log"
	authHttp "cloud-storage/internal/auth/http"
	"cloud-storage/internal/auth"
	"cloud-storage/internal/jwt"
	"cloud-storage/internal/auth/postgres"
	"cloud-storage/internal/config"
	"cloud-storage/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad("MAIN")	
	
	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	sessionRepo := postgres.NewSessionRepository(db)
	jwtManager := jwt.New(cfg.JWT)
	service := auth.NewService(userRepo, sessionRepo, jwtManager)
	handler := authHttp.NewHandler(service)
	router := gin.Default()

	authHttp.RegisterRoutes(router, handler)
	router.GET("/", func(c *gin.Context) {
		c.File("./web/test_auth.html")
	})

	addr := fmt.Sprintf("%s:%s",
		cfg.Server.Host,
		cfg.Server.Port,
	)

	log.Println("server started on", addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
