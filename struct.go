package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person = Person{"abby", 18}
	fmt.Println(p1.name, p1.age)
	var p2 Person = Person{age: 30, name: "aaron"}
	p2.name = "肥波"
	fmt.Println(p2.name, p2.age)
}
