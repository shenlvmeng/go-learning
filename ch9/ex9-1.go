package main

type Withdrawal struct {
	amount int
	ch     chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdrawal = make(chan Withdrawal)

func Deposite(amount int) { deposits <- amount }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawal <- Withdrawal{amount, ch}
	return <-ch
}
func Balance() int { return <-balances }

func monitor() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawal:
			if w.amount > balance {
				w.ch <- false
				continue
			}
			balance -= w.amount
			w.ch <- true
		case balances <- balance:
		}
	}
}

func init() {
	go monitor()
}

func main() {

}
