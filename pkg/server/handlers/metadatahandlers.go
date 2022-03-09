package handlers

import (
	"apiserver/pkg/server/api"
	"apiserver/pkg/store"
	"errors"
	"fmt"
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
		err := errors.New("Missing metadata id")
		logger.Error(err)
		// 400
		c.YAML(http.StatusBadRequest, api.NewErrorResponse(api.BadRequest, err.Error()))
		return
	}
	logger.Infof("metadata id: %s", id)

	// validate YAML payload
	if err := payload.Data.Validate(); err != nil {
		err = fmt.Errorf("Invalid YAML payload: %v", err)
		logger.Error(err)
		// 400
		c.YAML(http.StatusBadRequest, api.NewErrorResponse(api.BadRequest, err.Error()))
		return
	}

	if err := mm.Store.SetMetadata(id, payload.Data); err != nil {
		logger.Errorf("store SetMetadata error: %v", err)
		// 500
		c.YAML(http.StatusInternalServerError, api.NewErrorResponse(api.InternalServerError, "Cannot save metadata"))
		return
	}
	payload.ID = id
	// 201
	c.YAML(http.StatusCreated, payload)
}

func (mm *MetadataManager) GetMetadataHandler(c *gin.Context) {
	logger := mm.Logger.Sugar()
	id := c.Param(api.ParameterID)
	if id == "" {
		err := errors.New("Missing metadata id")
		logger.Error(err)
		// 400
		c.YAML(http.StatusBadRequest, api.NewErrorResponse(api.BadRequest, err.Error()))
		return
	}
	logger.Infof("metadata id: %s", id)

	md, err := mm.Store.GetMetadata(id)
	// 500
	if err != nil {
		logger.Errorf("store GetMetadata error: %v", err)
		c.YAML(http.StatusInternalServerError, api.NewErrorResponse(api.InternalServerError, "Cannot get metadata"))
		return
	}
	// 200
	if md != nil {
		c.YAML(http.StatusOK, api.NewPayload(id, md))
		return
	}
	// 404
	c.YAML(http.StatusNotFound, api.NewErrorResponse(api.NotFound, "Requested metadata not found"))
}

func (mm *MetadataManager) ListMetadataHandler(c *gin.Context) {
	logger := mm.Logger.Sugar()
	company := c.DefaultQuery(api.QueryCompany, "")
	if company == "" {
		err := errors.New("Only supports searching metadata by company name")
		logger.Error(err)
		// 400
		c.YAML(http.StatusBadRequest, api.NewErrorResponse(api.NotSupported, err.Error()))
		return
	}

	md, err := mm.Store.ListMedatadaByCompany(company)
	// 500
	if err != nil {
		logger.Errorf("store ListMedatadaByCompany error: %v", err)
		c.YAML(http.StatusInternalServerError, api.NewErrorResponse(api.InternalServerError, "Cannot list metadata"))
		return
	}

	c.YAML(http.StatusOK, md)
}
