package inter

import "fmt"

func (T *Cat) eat() {

	fmt.Println("I am eating")
}
func (t *Cat) look() {

	fmt.Println("I am looking")
}

type Eater interface {
	eat()
}
type Looker interface {
	look()
}

func TestDynamicAddInter() {

	var eater Eater
	var looker Looker
	cat := new(Cat)
	cat.Set("xixi")
	eater = cat
	eater.eat()
	looker = cat
	looker.look()
}
