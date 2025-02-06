package controller

import (
	"WMS/internal/domain"
	"WMS/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	service service.InventoryService
}

// Constructor function
func NewInventoryController(s service.InventoryService) *InventoryController {
	return &InventoryController{
		service: s,
	}
}

// Fetch inventory based on tenant_id, hub_id, and sku_id
func (c *InventoryController) FetchInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hubID, _ := strconv.Atoi(ctx.Query("hub_id"))
		skuID, _ := strconv.Atoi(ctx.Query("sku_id"))

		inventories, err := c.service.FetchInventory(ctx, hubID, skuID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
			return
		}
		ctx.JSON(http.StatusOK, inventories)
	}
}
func (c *InventoryController) UpdateInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inventory domain.Inventory
		if err := ctx.ShouldBindJSON(&inventory); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		err := c.service.UpdateInventory(ctx, inventory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
	}
}