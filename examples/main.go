package main

import (
	"fmt"
	"notifier"
)

func main() {
	myNotifier, err := notifier.New()
	if err != nil {
		panic(err)
	}

	// notifiers
	slack := myNotifier.NewSlackNotifier()

	// send message with slack notifier
	sendMessage(slack, "hello slack")
}

func sendMessage(notifier notifier.INotifier, message string) {
	fmt.Printf("\nSending message to %s...\n\n", notifier.Name())

	if err := notifier.Notify(message); err != nil {
		panic(err)
	}

	fmt.Printf("\nMessage send to %s!", notifier.Name())
}
