package main

// Struct that captures the structure of a receipt JSON
type Receipt struct {
	Retailer 			string	`json:"retailer" binding:"required"`
	PurchaseDate 	string 	`json:"purchaseDate" binding:"required"`
	PurchaseTime 	string 	`json:"purchaseTime" binding:"required"`
	Items 				[]Item	`json:"items" binding:"required"`
	Total 				float64	`json:"total,string" binding:"required"`
}
type Item struct {
	ShortDescription	string	`json:"shortDescription" binding:"required"`
	Price							float64	`json:"price,string" binding:"required"`
}
