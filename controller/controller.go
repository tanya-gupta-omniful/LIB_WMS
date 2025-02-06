package controller

import (
	"WMS/internal/domain"
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

func (c *Controller)GetSkus() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		skus := c.service.FetchSkus(ctx)
		ctx.JSON(http.StatusOK, skus)
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
		SkuID, err := strconv.Atoi(ctx.Param("id"))
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
func (c *Controller)GetHubByTenantId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		TenantID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ttenant ID"})
			return
		}
		hub, err := c.service.FetchHubByTenantId(ctx, TenantID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "hub not found for the tenant"})
			return
		}
		ctx.JSON(http.StatusOK, hub)
	}
}

// GetSkuBySellerID handles the API request to fetch SKU by seller_id
func (c *Controller) GetSkuBySellerID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		SellerID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Seller ID"})
			return
		}
		sku, err := c.service.FetchSkuBySellerID(ctx, SellerID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "SKU not found for the seller"})
			return
		}
		ctx.JSON(http.StatusOK, sku)
	}
}
// Create a new hub
func (c *Controller) CreateHub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var hub domain.Hub
		if err := ctx.ShouldBindJSON(&hub); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newHub, err := c.service.CreateHub(ctx, hub)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, newHub)
	}
}
// Create a new SKU
func (c *Controller) CreateSku() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sku domain.Sku
		if err := ctx.ShouldBindJSON(&sku); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newSku, err := c.service.CreateSku(ctx, sku)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, newSku)
	}
}
func (c *Controller) DeleteHub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hub ID"})
			return
		}
		err = c.service.DeleteHub(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Hub deleted successfully"})
	}
}
// Delete an SKU
func (c *Controller) DeleteSku() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skuID, err := strconv.Atoi(ctx.Param("sku_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
			return
		}
		err = c.service.DeleteSku(ctx, skuID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "SKU deleted successfully"})
	}
}