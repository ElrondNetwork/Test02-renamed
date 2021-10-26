package components

import "fmt"

type component struct {
	name   string
	points uint32
}

func NewComponent(name string, points uint32) *component {
	return &component{
		name:   name,
		points: points,
	}
}

func (c *component) PrintName() {
	fmt.Println(c.name)
}

func (c *component) PrintPoints() {
	fmt.Println(c.points)
}
