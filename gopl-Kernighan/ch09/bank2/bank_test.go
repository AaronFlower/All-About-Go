package bank

import (
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	var n sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			deposit(amount)
			n.Done()
		}(i)
	}

	n.Wait()

	if got, want := query(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want = %d", got, want)
	}
}
