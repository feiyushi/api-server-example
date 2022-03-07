package handlers

import (
	"apiserver/pkg/server/api"
	"apiserver/pkg/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MetadataManager struct {
	Store  store.MetadataStore
	Logger *zap.Logger
}

func (mm *MetadataManager) PutMetadataHandler(c *gin.Context) {
	logger := mm.Logger.Sugar()
	var payload api.Payload
	if err := c.ShouldBindYAML(&payload); err != nil {
		logger.Errorf("yaml binding error: %v", err)
		// 400
		c.YAML(http.StatusBadRequest, api.NewErrorResponse(api.BadRequest, "Invalid YAML payload"))
		return
	}

	id := c.Param(api.ParameterID)
	if id == "" {
		// 400
		c.YAML(http.StatusBadRequest, api.NewErrorResponse(api.BadRequest, "Missing metadata id"))
		return
	}
	logger.Infof("metadata id: %s", id)

	if err := mm.Store.SetMetadata(id, &payload); err != nil {
		logger.Errorf("store SetMetadata error: %v", err)
		// 500
		c.YAML(http.StatusInternalServerError, api.NewErrorResponse(api.InternalServerError, "Cannot save metadata"))
		return
	}
	// 201
	c.YAML(http.StatusCreated, payload)
}

func (mm *MetadataManager) GetMetadataHandler(c *gin.Context) {
	logger := mm.Logger.Sugar()
	id := c.Param(api.ParameterID)
	if id == "" {
		// 400
		c.YAML(http.StatusBadRequest, api.NewErrorResponse(api.BadRequest, "Missing metadata id"))
		return
	}
	logger.Infof("metadata id: %s", id)

	md, err := mm.Store.GetMetadata(id)
	// 500
	if err != nil {
		c.YAML(http.StatusInternalServerError, api.NewErrorResponse(api.InternalServerError, "Cannot get metadata"))
		return
	}
	// 200
	if md != nil {
		c.YAML(http.StatusOK, md)
		return
	}
	// 404
	c.YAML(http.StatusNotFound, api.NewErrorResponse(api.NotFound, "Requested metadata not found"))
}

func (mm *MetadataManager) ListMetadataHandler(c *gin.Context) {

}
