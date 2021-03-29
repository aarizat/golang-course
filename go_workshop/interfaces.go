package main

import "fmt"


type animal interface {
	move() string
}

type dog struct {}
type fish struct {}
type bird struct {}

func (dog) move() string {
	return "Dog is walking"
}

func (fish) move() string {
	return "Fish is swimming"
}

func (bird) move() string {
	return "Bird is flying"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}


func main() {
	d := dog{}
	f := fish{}
	b := bird{}

	moveAnimal(d)
	moveAnimal(f)
	moveAnimal(b)
}