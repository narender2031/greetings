package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf(rendomGreetings(), name)

	return message, nil
}

// Add greetings for multiple names

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[message] = message
	}

	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func rendomGreetings() string {
	greetingMessages := []string{
		"Hi %v! welcome to the go group",
		"%v, Welcome to the team ",
		"Welcome! %v",
	}

	return greetingMessages[rand.Intn(len(greetingMessages))]
}
