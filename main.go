package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create new router with default middleware (logger and recovery)
	r := gin.Default()

	// Run and listen on port 0.0.0.0:8080
	r.Run() 
}
