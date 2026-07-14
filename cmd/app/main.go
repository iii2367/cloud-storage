package main

import (
	"cloud-storage/internal/auth/repository/postgres"
	"cloud-storage/internal/auth/service"
	authHttp "cloud-storage/internal/auth/transport/http"
	"cloud-storage/internal/config"
	"cloud-storage/internal/database"
	"cloud-storage/internal/jwt"
	"cloud-storage/internal/web"
	"fmt"
	"log"

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
	service := service.NewService(userRepo, sessionRepo, jwtManager)

	handler := authHttp.NewHandler(service)
	middleware := authHttp.NewMiddleware(jwtManager)

	router := gin.Default()
	
	router.LoadHTMLGlob("web/pages/*")
	router.Static("/static", "./web/static/css")	

	authHttp.RegisterRoutes(router, handler, middleware)
	
	webHandler := web.NewHandler()
	web.RegisterRoutes(router, webHandler)

	router.GET("/test_auth", func(c *gin.Context) {
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
