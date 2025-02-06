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

func (c *Controller) GetSkuByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get sku_id from the URL parameters
		SkuID, err := strconv.Atoi(ctx.Param("sku_id"))
		if err != nil {
			// Handle invalid SKU ID format
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
			return
		}

		// Fetch the SKU based on the SKU ID
		sku, err := c.service.FetchSkuByID(ctx, SkuID)
		if err != nil {
			// Handle error if no SKU is found
			ctx.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
			return
		}

		// Return the SKU details if found
		ctx.JSON(http.StatusOK, sku)
	}
}

// GetSkuBySellerID handles the API request to fetch SKU by seller_id
func (c *Controller) GetSkuBySellerID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get seller_id from the URL parameters
		SellerID, err := strconv.Atoi(ctx.Param("seller_id"))
		if err != nil {
			// Handle invalid Seller ID format
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Seller ID"})
			return
		}

		// Fetch the SKU based on the Seller ID
		sku, err := c.service.FetchSkuBySellerID(ctx, SellerID)
		if err != nil {
			// Handle error if no SKU is found for the given seller_id
			ctx.JSON(http.StatusNotFound, gin.H{"error": "SKU not found for the seller"})
			return
		}

		// Return the SKU details if found
		ctx.JSON(http.StatusOK, sku)
	}
}
