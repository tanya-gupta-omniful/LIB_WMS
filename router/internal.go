package router

import (
	"WMS/controller"
	pkg "WMS/db"
	"WMS/repo"
	"WMS/service"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
)

func InternalRoutes(ctx context.Context, s *http.Server) (err error) {
	rtr := s.Engine.Group("/api/v1")

	// todo use go wire if needed
	newRepository := repo.NewRepository(pkg.GetCluster().DbCluster)
	newService := service.NewService(newRepository)
	controller := controller.NewController(newService)

	// make apis for it
	rtr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "mst"})
	}) 
    rtr.POST("/hub",controller.CreateHub())
	rtr.GET("/hub", controller.GetHubs())
	rtr.GET("/hub/:id", controller.GetHubByID())
	rtr.GET("/hub/tenant/:id", controller.GetHubByTenantId())
	rtr.DELETE("/hub/:id",controller.DeleteHub())
	rtr.POST("/sku",controller.CreateSku())
	rtr.GET("/sku", controller.GetSkus())
	rtr.GET("/sku/:id", controller.GetSkuByID())
	rtr.DELETE("/sku/:id",controller.DeleteSku())

	// Define GET route to fetch SKUs by Seller ID
	rtr.GET("/sku/seller/:id", controller.GetSkuBySellerID())

	return
}