// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to understand method sets.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func main() {

	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}
	m := user{"Arshpreet", "arsh840@gmail.com"}
	// Values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value.
	u.notify()
	m.notify()
}
