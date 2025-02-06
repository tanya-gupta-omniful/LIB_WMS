package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"WMS/services"

	"github.com/gin-gonic/gin"
)

// Controller - Struct for handling hub services
type Controller struct {
	hubService services.HubService
}

var ctrl *Controller
var ctrlOnce sync.Once

// NewController - Singleton pattern for Controller initialization
func NewController(svc services.HubService) *Controller {
	ctrlOnce.Do(func() {
		ctrl = &Controller{
			hubService: svc,
		}
	})
	return ctrl
}

// HubResponse - Defines the response structure for hub requests
type HubResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// GetHub - Fetch details of a single hub
func (tc *Controller) GetHub(c *gin.Context) {
	hubID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
		return
	}

	hub, err := tc.hubService.GetHubDetails(hubID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		return
	}

	response := HubResponse{
		ID:       hub.ID,
		Name:     hub.Name,
		Location: hub.Location,
	}

	c.JSON(http.StatusOK, response)
}

// CreateHub - Create a new warehouse hub
func (tc *Controller) CreateHub(c *gin.Context) {
	var hub models.Hub
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdHub, err := tc.hubService.CreateHub(hub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := HubResponse{
		ID:       createdHub.ID,
		Name:     createdHub.Name,
		Location: createdHub.Location,
	}

	c.JSON(http.StatusCreated, response)
}
