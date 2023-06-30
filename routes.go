package main

import (
	"github.com/gin-gonic/gin"
)

/*
 * Define two routes for this service.
 * @param r -> A router created with the Gin framework.
 */
func setupRoutes(r *gin.Engine) {
	r.POST("/receipts/process", processReceiptHandler)
	r.GET("/receipts/:id/points", getPointsHandler)
}
