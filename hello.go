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

type admin struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println(u.name)
	fmt.Println(u.email)
}

func (a admin) notify() {
	fmt.Println("不允许查看管理员通知信息")
}

func main() {
	fmt.Println("hello,world!")
	bill := user{"Bill", "Bill@ink.com"}
	bill.notify()
}
