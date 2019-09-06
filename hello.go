package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println(u.name)
	fmt.Println(u.email)
}

func main() {
	fmt.Println("hello,world!")
	bill := user{"Bill", "Bill@ink.com"}
	bill.notify()
}
