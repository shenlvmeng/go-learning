package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func scan(r io.Reader, lines chan<- string) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		lines <- input.Text()
	}
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
		c.Close()
	}()
	lines := make(chan string)
	go scan(c, lines)
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case line := <-lines:
			timer.Reset(10 * time.Second)
			wg.Add(1)
			go func(text string) {
				echo(c, text, 1*time.Second)
				wg.Done()
			}(line)
		case <-timer.C:
			return
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
