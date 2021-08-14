package main

import (
	"fmt"

	"github.com/suleymantaspinar/hello-go/functions"
	"github.com/suleymantaspinar/hello-go/variables"
)

func main() {
	fmt.Println("Hey Dude")
	variables.Run()
	functions.Run(15, 15)
}
