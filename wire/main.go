package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Message string
type Greeter struct {
	Message Message
	Grumpy bool
}
type Event struct {
	Greeter Greeter
}
// NewMessage Message的构造函数
func NewMessage(phrase string) Message {
	return Message(phrase)
}
// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	var grumy bool
	if time.Now().Unix()%2 == 0 {
		grumy = true
	}
	return Greeter{Message: m, Grumpy: grumy}
}
// NewEvent Event构造函数
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go Away")
	}
	return g.Message
}

func main() {
	e, err := InitializeEvent("Hello")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
