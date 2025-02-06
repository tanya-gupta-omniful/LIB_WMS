package controllers

import (
	"net/http"
	"strconv"
	"sync"

	 "WMS/internal/domain"

	"github.com/gin-gonic/gin"
)

// InventoryController - Struct for handling inventory-related requests
type InventoryController struct {
	inventoryService domain.InventoryService
}

var invCtrl *InventoryController
var invCtrlOnce sync.Once

// NewInventoryController - Singleton pattern for InventoryController initialization
func NewInventoryController(invSvc domain.InventoryService) *InventoryController {
	invCtrlOnce.Do(func() {
		invCtrl = &InventoryController{
			inventoryService: invSvc,
		}
	})
	return invCtrl
}

// GetInventory - Fetch inventory details for a given seller and hub
func (ic *InventoryController) GetInventory(c *gin.Context) {
	sellerID, err := strconv.ParseUint(c.Query("seller_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	hubID, err := strconv.ParseUint(c.Query("hub_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
		return
	}

	inventory, err := ic.inventoryService.GetInventory(c, sellerID, hubID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// EditInventory - Update or create inventory (Upsert)
func (ic *InventoryController) EditInventory(c *gin.Context) {
	var req struct {
		SKUCode  string `json:"sku_code" binding:"required"`
		HubID    uint64 `json:"hub_id" binding:"required"`
		Quantity int    `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ic.inventoryService.UpdateInventory(c, req.SKUCode, req.HubID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}
