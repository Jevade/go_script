package car

import (
	"fmt"
)

// func main(){
// 	car.Stop()
// 	car.start()
// 	car.stop()
// }

type Engine interface{
	stop()
	start()
}

func (T *Car)stop(){
	fmt.Println("car stop")
}

func (T *Car)start(){
	fmt.Println("car start")
}

type Car struct{
	wheel int
    Engine	
}

func (T * Car)numberOfWheel()int{
	return T.wheel
}

type Mecideris struct{
	Car
	logo string
}

func NewMecideris(wheel int,logo string)(mecideris *Mecideris){
	mecideris = new(Mecideris)
	mecideris.wheel = wheel
	mecideris.logo = logo
	return 
}
func (T *Mecideris)NumberOfWheel()int{
	return T.numberOfWheel()
}

func (T *Mecideris)Start(){
	fmt.Println("I'am :",T.logo)
	T.Car.start()
}

func (T *Mecideris)Stop(){
	fmt.Println("I'am :",T.logo)
	T.Car.stop()
}