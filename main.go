package main

import (
	"fmt"
	"go_leaning/animals"
)

func main() {
	fmt.Println("Hello, World")

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
