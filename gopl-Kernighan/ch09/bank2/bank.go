package bank

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance.
	balance int
)

func deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func query() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
