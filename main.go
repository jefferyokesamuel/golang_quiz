package main

import (
	"fmt"
	"log"
	// "os"
	// "io"
	"github.com/Delta456/box-cli-maker/v2"
	"github.com/dixonwille/wmenu/v5"
)

func main() {

	var right int
	
	Box := box.New(box.Config{
		Px:         20,
		Py:         10,
		Type:       "Hidden",
		Color:      "Yellow",
		TitleColor: "Red",
	})
	Box.Print("Welcome🔥🚀", "This is a Golang Quiz project get ready to answer some trivia")

	q1 := wmenu.NewMenu("What's the capital of Germany?")
	q1.Action(func (opts []wmenu.Opt) error {
		fmt.Printf("Your answer " + opts[0].Text + "has been locked in\n")
		if opts[0].Text == "Berlin" {
			right = right + 1
		}
		return nil
	})
	q1.Option("Berlin", nil, false, nil)
	q1.Option("Amsterdam", nil, false, func(wmenu.Opt) error {
		fmt.Println("You can't be serious")
		return nil
	})
	q1.Option("Munich", nil, false, nil )
	err := q1.Run()
	if err != nil {
		log.Fatal(err)
	}


	q2 := wmenu.NewMenu("What's the first Cryptocurrency?")
	q2.Action(func (opts []wmenu.Opt) error {
		fmt.Printf("Your answer " + opts[0].Text + "has been locked in\n")
		if opts[0].Text == "Bitcoin" {
			right = right + 1
		}
		return nil
	})
	q2.Option("Bitcoin", nil, true, nil)
	q2.Option("Ethereum", nil, false, nil)
	q2.Option("Atom", nil, false, nil)
	err2 := q2.Run()
	if err2 != nil{
	  log.Fatal(err)
	}

	fmt.Printf("Your score is %v out of 2\n", right)
}
