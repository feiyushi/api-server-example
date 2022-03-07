package main

import (
	"apiserver/pkg/server/api"
	"apiserver/pkg/server/handlers"
	"apiserver/pkg/store"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	setUpRouter()
}

func setUpRouter() {
	router := gin.New()
	logger, _ := zap.NewProduction()

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	// initialize metadata store
	store := store.NewInMemoryMetadataStore()
	metadataManager := &handlers.MetadataManager{Store: store, Logger: logger}
	// PUT metadata
	router.PUT(api.MetadataByIdRoute, metadataManager.PutMetadataHandler)

	// GET metadata
	router.GET(api.MetadataByIdRoute, metadataManager.GetMetadataHandler)

	// List metadata
	router.GET(api.MetadataRoute, metadataManager.ListMetadataHandler)

	// listen on 8080 port
	router.Run(":8080")
}
