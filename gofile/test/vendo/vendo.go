package vendo

import (
	"fmt"
)

type Base struct{

}

func (Base) Magic(){
	fmt.Println("From Base magic")
}

func (self Base) MoreMagic(){
	self.Magic()
	self.Magic()
}

type Vendo struct{
	Base
}

func (Vendo) Magic(){
	fmt.Println("From Vendo magic")
}
