package main

import "fmt"

type animal interface {
	move() string
}

type dog struct {

}

type fish struct {

}

type bird struct {

}

func (dog) move() string {
	return "I'm dog and i walk"
}

func (fish) move() string {
	return "I'm fish and i'm swimming"
}

func (bird) move() string {
	return "I'm bird and i'm fling"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	d := dog{}
	moveAnimal(d)
	f := fish{}
	moveAnimal(f)
	b := bird{}
	moveAnimal(b)
}
