package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Set id while testing processReceiptHander(), and use it to test getPointsHandler().
var valid_id string

/*
 * Initialize global Receipt map called "ReceiptMap" before tests run.
 */
func TestMain(m *testing.M) {
	ReceiptMap.Init()
	os.Exit(m.Run())
}

/*
 * Tests processReceiptHander() with a valid request body.
 */
func TestProcessReceiptHandlerValid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create valid JSON body
	body := Receipt{
		Retailer: "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item {
			{ 
				ShortDescription: "Mountain Dew 12PK",
				Price: 6.49,
			},
			{ 
				ShortDescription: "Emils Cheese Pizza",
				Price: 12.25,
			},
			{ 
				ShortDescription: "Knorr Creamy Chicken",
				Price: 1.26,
			},
			{ 
				ShortDescription: "Doritos Nacho Cheese",
				Price: 3.35,
			},
			{ 
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price: 12.00,
			},
		},
		Total: 35.35,
	}
	bodyBytes, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(bodyBytes)

	// Create test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/some-path/123", bodyReader)
	c.Request.Header.Set("Content-Type", "application/json")

	processReceiptHandler(c)

	// Check the response code (should be 201)
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %v, got %v", http.StatusCreated, w.Code)
	}

	// Parse the JSON response body into a map
	bodyString := w.Body.String()
	var result map[string]string
	err := json.Unmarshal([]byte(bodyString), &result)
	if err != nil {
		t.Fatal("Error parsing JSON:", err)
	}

	// Check if id field exists in the response
	id, exists := result["id"]
	if !exists {
		t.Fatal("Response body does not contain an id field")
	}

	// Check if id field is a valid UUID
	_, err = uuid.Parse(id)
	if err != nil {
		t.Fatal("ID is not a valid UUID", err)
	}

	valid_id = id
}

/*
 * Tests processReceiptHandler with no request body.
 */
func TestProcessReceiptHandlerNoBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/some-path/123", nil)

	processReceiptHandler(c)

	// Check the response code (should be 400)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %v, got %v", http.StatusBadRequest, w.Code)
	}
}

/*
 * Tests processReceiptHander() with a purchaseDate with an invalid format.
 */
func TestProcessReceiptHandlerInvalidDate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create valid JSON body
	body := Receipt{
		Retailer: "Target",
		PurchaseDate: "20-01-01",
		PurchaseTime: "13:01",
		Items: []Item {
			{ 
				ShortDescription: "Mountain Dew 12PK",
				Price: 6.49,
			},
			{ 
				ShortDescription: "Emils Cheese Pizza",
				Price: 12.25,
			},
			{ 
				ShortDescription: "Knorr Creamy Chicken",
				Price: 1.26,
			},
			{ 
				ShortDescription: "Doritos Nacho Cheese",
				Price: 3.35,
			},
			{ 
				ShortDescription: "  Klarbrun 12-PK 12 FL OZ  ",
				Price: 12.00,
			},
		},
		Total: 35.35,
	}
	bodyBytes, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(bodyBytes)

	// Create test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/some-path/123", bodyReader)
	c.Request.Header.Set("Content-Type", "application/json")

	processReceiptHandler(c)

	// Check the response code (should be 400)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %v, got %v", http.StatusBadRequest, w.Code)
	}
}

/*
 * Tests processReceiptHander() with a purchaseTime with an invalid format.
 */
func TestProcessReceiptHandlerInvalidTime(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create valid JSON body
	body := Receipt{
		Retailer: "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "1:1",
		Items: []Item {
			{ 
				ShortDescription: "Mountain Dew 12PK",
				Price: 6.49,
			},
			{ 
				ShortDescription: "Emils Cheese Pizza",
				Price: 12.25,
			},
			{ 
				ShortDescription: "Knorr Creamy Chicken",
				Price: 1.26,
			},
			{ 
				ShortDescription: "Doritos Nacho Cheese",
				Price: 3.35,
			},
			{ 
				ShortDescription: "  Klarbrun 12-PK 12 FL OZ  ",
				Price: 12.00,
			},
		},
		Total: 35.35,
	}
	bodyBytes, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(bodyBytes)

	// Create test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/some-path/123", bodyReader)
	c.Request.Header.Set("Content-Type", "application/json")

	processReceiptHandler(c)

	// Check the response code (should be 400)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %v, got %v", http.StatusBadRequest, w.Code)
	}
}

/*
 * Test getPointsHandler with a valid id.
 */
func TestGetPointsHandlerValid(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a router so we can test with the id from the path
	r := gin.Default()
	r.GET("/receipt/:id/points", getPointsHandler)

	// Create request to the route
	validPath := "/receipt/" + valid_id + "/points"
	req, _ := http.NewRequest(http.MethodGet, validPath, nil)

	// Server request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response code (should be 200)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}

	// Parse the JSON response body into a map
	bodyString := w.Body.String()
	var result map[string]int
	err := json.Unmarshal([]byte(bodyString), &result)
	if err != nil {
		t.Fatal("Error parsing JSON:", err)
	}

	// Check if points field exists in the response
	points, exists := result["points"]
	if !exists {
		t.Fatal("Response body does not contain a points field")
	}

	if points != 28 {
		t.Fatalf("Expected points value 28, got %v", points)
	}
}

/*
 * Tests getPointsHandler with an invalid id.
 */
func TestGetPointsHandlerNoBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/some-path/123", nil)

	getPointsHandler(c)

	// Check the response code (should be 400)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %v, got %v", http.StatusBadRequest, w.Code)
	}
}