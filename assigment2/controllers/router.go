package controllers

import (

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/api")
	routes.POST("/order", h.AddOrder)
	routes.GET("/order", h.GetOrders)
	routes.GET("/order/:id", h.GetOrderById)
	routes.PUT("/order/:id", h.UpdateOrder)
	routes.DELETE("/order/:id", h.DeleteOrder)
}
