package main

import (
	authRepoPostgres "cloud-storage/internal/auth/repository/postgres"
	authService "cloud-storage/internal/auth/service"
	authHttp "cloud-storage/internal/auth/transport/http"

	/*
	storageRepoPostgres "cloud-storage/internal/storage/repository/postgres"
	storageService "cloud-storage/internal/storage/service"
	storageHttp "cloud-storage/internal/storage/transport/http"
	*/

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
	jwtManager := jwt.New(cfg.JWT)

	userRepo := authRepoPostgres.NewUserRepository(db)
	sessionRepo := authRepoPostgres.NewSessionRepository(db)
	authService := authService.NewService(userRepo, sessionRepo, jwtManager)
	authHandler := authHttp.NewHandler(authService)
	authMiddleware := authHttp.NewMiddleware(jwtManager)

	/*
	treeNodeRepo := storageRepoPostgres.NewTreeNodeRepository(db)
	storageService := storageService.NewService(treeNodeRepo, jwtManager)
	storageHandler := storageHttp.NewHandler(storageService)
	storageMiddleware := storageHttp.NewMiddleware(jwtManager)
	*/

	router := gin.Default()
	
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/css", "./web/static/css")	
	router.Static("/js", "./web/static/js")	

	authHttp.RegisterRoutes(router, authHandler, authMiddleware)
	//storageHttp.RegisterRoutes(router, storageHandler, storageMiddleware)
	
	webHandler := web.NewHandler()
	web.RegisterRoutes(router, webHandler)

	router.GET("/test_auth", func(c *gin.Context) {
		c.File("./web/templates/test_auth_api.html")
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
