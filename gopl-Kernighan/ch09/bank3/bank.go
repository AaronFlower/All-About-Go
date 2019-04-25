package bank

import "sync"

// By convention, the variables guarded by a mutex are declared immediately after the declaration
// of the mutex itself.
var (
	mu      sync.Mutex // guards balance
	balance int
)

func deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func query() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
