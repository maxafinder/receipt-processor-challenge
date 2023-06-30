package main

//"errors"
import (
	"errors"
	"math"
	"strings"
	"sync"
	"time"
	"unicode"
)

/*
 * This function calculates the points for a receipt based on these rules:
 *
 *	1.	One point for every alphanumeric character in the retailer name.
 *	2.  50 points if the total is a round collar amount with no cents.
 *	3.	25 points if the toal is a multiple of 0.25.
 *	4.	5 points for every two items on the receipt.
 *	5.	If the trimmed length of the item description is a multiple of 3, multiply
 *			the price by 0.2 and round up to the nearest integer. The result is the
 *			number of points earned.
 *	6.	6 points if the day in the purchase date is odd.
 *	7.	10 points if the time of purchase is after 2:00pm and before 4:00pm.
 *
 * @param r -> the receipt we are calculating the point for.
 * @return -> the score of r
 */
func calculateReceiptPoints(r Receipt) (int, error) {
	points := 0

	// Rule 1 
	alphNumCount := 0
	for _, ch := range r.Retailer {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			alphNumCount++
		}
	}
	points += alphNumCount

	// Rule 2
	if r.Total == math.Floor(r.Total) {
		points += 50
	}

	// Rule 3
	if math.Mod(r.Total, 0.25) == 0 {
		points += 25
	}

	// Rule 4
	numItemPairs := len(r.Items) / 2
	points += 5 * numItemPairs

	// Rule 5
	for _, item := range r.Items {
		trimmedItemDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedItemDesc) % 3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	// Rule 6
	date, dateErr := time.Parse("2006-01-02", r.PurchaseDate)
	if dateErr != nil {
		return 0, errors.New("problem parsing purchaseDate data for this receipt")
	}
	if date.Day() % 2 != 0 {
		points += 6
	}

	// Rule 7
	time, timeErr := time.Parse("15:04", r.PurchaseTime)
	if timeErr != nil {
		return 0, errors.New("problem parsing purchaseTime data for this receipt")
	}
	if time.Hour() >= 14 && time.Hour() < 16 {
		points += 10
	}

	return points, nil
}

/*
 * Map structure that provides mutual exclusion using mutex locks.
 * The Gin framework handles requests concurrently so the shared map resource
 * needs to be mutually explusive.
 */
type SafeMap struct {
	smap map[string]Receipt // empty interface can hold any type
	mut sync.Mutex
}

// Initialize SafeMap
func (m *SafeMap) Init() {
	m.smap = make(map[string]Receipt)
}

// Set key/value in SafeMap
func (m *SafeMap) SafeSet(key string, value Receipt) {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.smap[key] = value
}

// Get value from SafeMap given a key
func (m *SafeMap) SafeGet(key string) (Receipt, bool) {
	m.mut.Lock()
	defer m.mut.Unlock()
	val, exists := m.smap[key]
	return val, exists
}
