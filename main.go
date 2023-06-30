package main

import (
	"github.com/gin-gonic/gin"
)

// Global map to store Receipts by ID
var ReceiptMap SafeMap

func main() {
	ReceiptMap.Init()

	// Create new router with default middleware (logger and recovery)
	r := gin.Default()
	setupRoutes(r)

	// Run and listen on port from config.go
	r.Run(":" + Port) 
}
