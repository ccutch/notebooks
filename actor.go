package main

import "log"

type Increment int
type Hello string
type Actor struct {
	Mailbox chan interface{}
	Value   int
}

func (a *Actor) Run() {
	for {
		m := <-a.Mailbox
		log.Println("Message recieved")
		switch msg := m.(type) {
		case Increment:
			a.Value += int(msg)
		case Hello:
			log.Println("Hello " + string(msg) + "!")
		}
	}
}

func main() {
	mailbox := make(chan interface{})
	a := &Actor{mailbox, 10}

	go a.Run()

	//a.Mailbox <- Increment(15)
	a.Mailbox <- Hello("Ed")
	log.Println("Done")
}

