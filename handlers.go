package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
 * Function that handles the POST request for route "/receipts/process".
 * 1. Validates the request body and give a 400 response if invalid.
 * 2. Generates an ID for the object and saves it in a map (local storage).
 * 3. If successful, give a 201 response and return a JSON object with ID.
 * @param c -> A Gin struct containing information to handle the HTTP request.
 */
func processReceiptHandler(c *gin.Context) {
	var receipt Receipt
	err := c.ShouldBindJSON(&receipt)
	print(err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid receipt received",
		})
		return
	}

	// Check that purchaseDate and purchaseTime are in the right format
	_, dateErr := time.Parse("2006-01-02", receipt.PurchaseDate)
	_, timeErr := time.Parse("15:04", receipt.PurchaseTime)
	if dateErr != nil {
		c.JSON(400, gin.H{
			"error": "Date is in the wrong format",
		})
		return
	}
	if timeErr != nil {
		c.JSON(400, gin.H{
			"error": "Time is in the wrong format",
		})
		return
	}

	newId := uuid.New()
	ReceiptMap.SafeSet(newId.String(), receipt)

	c.JSON(201, gin.H{
		"id": newId,
	})
}

/*
 * Function that handles the GET request for route "/receipts/{id}/points".
 * 1. Validate the receipt with the ID exists, and give 400 if it doesn't.
 * 2. Calculate the number of points for that receipt.
 * 3. If successful, give a 200 response and return a JSON object with points.
 * @param c -> A Gin struct containing information to handle the HTTP request.
 */
func getPointsHandler(c *gin.Context) {
	id := c.Param("id")
	receipt, exists := ReceiptMap.SafeGet(id)
	if !exists {
		errMsg := "Couldn't find receipt with id=" + id
		c.JSON(400, gin.H{
			"error": errMsg,
		})
		return
	}

	points, err := calculateReceiptPoints(receipt)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"points": points, 
	})
}