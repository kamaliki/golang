package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(1)

	//a slice of names
	names := []string{"Clve", "Vic", "Bert"}

	//request a greeting message
	//message, err := greetings.Hello("Clive")
	// if an error was returned, print it to the console and
	// exit the program

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message
	// to the console
	fmt.Println(messages)
}