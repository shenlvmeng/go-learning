package main

type client struct {
	Name string
	Out  chan<- string // 只写信道
}

var (
	entering = make(chan client)
	messages = make(chan string)
	leaving  = make(chan client)
)

func broadcaster() {
	clients := make(map[client]bool) // 用户session维护
	for {
		select {
		case cli := <-entering:
			cli.Out <- "Current users: "
			for c := range clients {
				cli.Out <- c.Name
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		case msg := <-messages:
			for cli := range clients {
				cli.Out <- msg
			}
		}
	}
}
