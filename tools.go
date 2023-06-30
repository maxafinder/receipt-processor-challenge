package main

//"errors"
import (
	"sync"
)

// TODO: create function to calculate points for a receipt
/*
 *
 */
func calculateReceiptPoints(r Receipt) int {
	return 0
}

// TODO: create functions for dealing with time 
/*
 *
 */


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
