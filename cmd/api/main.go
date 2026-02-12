package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"ig4llc.com/internal/domain/events"
	"ig4llc.com/internal/infrastructure/db"
	httphandler "ig4llc.com/internal/infrastructure/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal("DB init failed:", err)
	}
	defer db.CloseDB()

	repo := db.NewEventRepo(db.DB)
	svc := events.NewService(repo)
	handler := httphandler.NewEventHandler(svc)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //add Swagger middleware
	handler.RegisterRoutes(r)
	r.Run(":8081")
}
