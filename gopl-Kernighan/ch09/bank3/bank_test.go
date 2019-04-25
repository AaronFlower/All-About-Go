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

	if want, got := (1+1000)*1000/2, query(); want != got {
		t.Errorf("want %d but got %d", want, got)
	}
}
