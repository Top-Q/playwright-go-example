package unit

import (
	"fmt"
	"testing"
)

type person struct {
	id   int
	name string
}

func (p *person) changeName(newName string) {
	p.name = newName
}

func personChanger(p *person) {
	p.changeName("Itai")
	fmt.Println(p)
}

func TestMyPointers(t *testing.T) {
	p := person{1, "Yossi"}
	personChanger(&p)
	fmt.Println(p)

}
