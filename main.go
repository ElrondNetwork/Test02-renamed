package main

import "github.com/ElrondNetwork/Test02/components"

func main() {
	c := components.NewComponent("test", 100)
	c.PrintName()
	c.PrintPoints()
}
