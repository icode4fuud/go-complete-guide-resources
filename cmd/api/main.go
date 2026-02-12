package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"ig4llc.com/internal/domain/events"
	"ig4llc.com/internal/infrastructure/db"
	httphandler "ig4llc.com/internal/infrastructure/http"
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
	handler.RegisterRoutes(r)
	r.Run(":8081")
}
