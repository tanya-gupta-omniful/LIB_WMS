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
	maincontroller := controller.NewController(newService)
	
	inventoryRepository := repo.NewInventoryRepository(pkg.GetCluster().DbCluster)
	inventoryService := service.NewInventoryService(inventoryRepository)
	inventoryController := controller.NewInventoryController(inventoryService)


	// make apis for it
	rtr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "mst"})
	}) 
    rtr.POST("/hub",maincontroller.CreateHub())
	rtr.GET("/hub", maincontroller.GetHubs())
	rtr.GET("/hub/:id", maincontroller.GetHubByID())
	rtr.GET("/hub/tenant/:id", maincontroller.GetHubByTenantId())
	rtr.DELETE("/hub/:id",maincontroller.DeleteHub())
	rtr.POST("/sku",maincontroller.CreateSku())
	rtr.GET("/sku", maincontroller.GetSkus())
	rtr.GET("/sku/:id", maincontroller.GetSkuByID())
	rtr.DELETE("/sku/:id",maincontroller.DeleteSku())
	rtr.GET("/sku/seller/:id", maincontroller.GetSkuBySellerID())
	rtr.GET("/inventory", inventoryController.FetchInventory()) 
    rtr.PUT("/inventory/", inventoryController.UpdateInventory())
	return
}