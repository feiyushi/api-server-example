package main

import (
	"apiserver/pkg/server/api"
	"apiserver/pkg/server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	setUpRouter()
}

func setUpRouter() {
	router := gin.Default()

	metadataManager := &handlers.MetadataManager{}
	// PUT metadata
	router.PUT(api.MetadataByIdRoute, metadataManager.PutMetadataHandler)

	// GET metadata
	router.GET(api.MetadataByIdRoute, metadataManager.GetMetadataHandler)

	// List metadata
	router.GET(api.MetadataRoute, metadataManager.ListMetadataHandler)

	// listen on 8080 port
	router.Run(":8080")
}
