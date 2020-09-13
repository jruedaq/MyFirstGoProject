package main

import "fmt"

type dog struct {

}

type fish struct {

}

type bird struct {

}

func (dog) walk() string {
	return "I'm dog and i walk"
}

func (fish) swim() string {
	return "I'm fish and i'm swimming"
}

func (bird) fly() string {
	return "I'm bird and i'm fling"
}

func moveDog(d dog) {
	fmt.Println(d.walk())
}

func moveFish(f fish) {
	fmt.Println(f.swim())
}

func moveBird(b bird) {
	fmt.Println(b.fly())
}

func main() {
	d := dog{}
	moveDog(d)
	f := fish{}
	moveFish(f)
	b := bird{}
	moveBird(b)
}
