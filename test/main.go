package main

import "test/storage/list"

func main() {
	l := list.List{}
	l.Add(10)
	l.Add(20)
	l.Add(30)
	l.Print()
}

