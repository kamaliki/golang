package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//randomFormat returns one of a set of greeting messages. The returned
//message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

//Hellos returns a map that associates each of the named people
//with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	// Assign the message to the map using the name as the key.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			// Return the error to the caller and nil
			// map.
			return nil, err
		}
		// Assign the message to the map using the name
		// as the key.
		messages[name] = message
	}
	return messages, nil
}

