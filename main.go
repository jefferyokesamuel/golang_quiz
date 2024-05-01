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

	q1 := wmenu.NewMenu("What is your favorite food?")
	q1.Action(func (opts []wmenu.Opt) error {
		fmt.Printf(opts[0].Text + " is your favorite food.\n")
		if opts[0].Text == "Pizza" {
			right = right + 1
		}
		return nil
	})
	q1.Option("Pizza", nil, true, nil)
	q1.Option("Ice Cream", nil, false, nil)
	q1.Option("Tacos", nil, false, func(opt wmenu.Opt) error {
		fmt.Printf("Tacos are great")
		return nil
	})
	err := q1.Run()
	if err != nil {
		log.Fatal(err)
	}


	q2 := wmenu.NewMenu("What is your favorite food?")
	q2.Action(func (opts []wmenu.Opt) error {
		fmt.Printf(opts[0].Text + " is your favorite food.\n")
		if opts[0].Text == "Pizza" {
			right = right + 1
		}
		return nil
	})
	q2.Option("Pizza", nil, true, nil)
	q2.Option("Ice Cream", nil, false, nil)
	q2.Option("Tacos", nil, false, func(opt wmenu.Opt) error {
	  fmt.Printf("Tacos are great")
	  return nil
	})
	err2 := q2.Run()
	if err2 != nil{
	  log.Fatal(err)
	}

	fmt.Printf("Your score is %v out of 2\n", right)
}
