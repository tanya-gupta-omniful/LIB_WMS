package controller

import (
	"WMS/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) GetHubs() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		hubs := c.service.FetchHubs(ctx)
		//if err != nil {
		//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hubs"})
		//	return
		//}

		ctx.JSON(http.StatusOK, hubs)

	}
}

func (c *Controller) GetHubByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hubID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
			return
		}

		hub, err := c.service.FetchHubByID(ctx, hubID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
			return
		}

		ctx.JSON(http.StatusOK, hub)
	}
}

func (c *Controller) GetHubByTenantID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		TenantID, err := strconv.Atoi(ctx.Param("tenant_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Tenant ID"})
			return
		}

		hub, err := c.service.FetchHubByTenantID(ctx, TenantID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Hub not found for this tenant"})
			return
		}

		ctx.JSON(http.StatusOK, hub)
		
	}
}