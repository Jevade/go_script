package car

import (
	"fmt"
)

//Engine is type can stop and start
type Engine interface {
	stop()
	start()
}

func (T *Car) stop() {
	fmt.Println("car stop")
}

func (T *Car) start() {
	fmt.Println("car start")
}

//Car is a struct with wheel int and Engine
type Car struct {
	wheel      int
	Model      string
	Manufactor string
	BuildYear  int
	Engine
}

func (T *Car) numberOfWheel() int {
	return T.wheel
}

//ExpensiveCar is dear

//Mecideris is the type of car
type Mecideris struct {
	Car
	logo string
}

//BMW is dear
type BMW struct {
	Car
	Price float64
}

//NewMecideris will create a new car with given params
func NewMecideris(wheel int, logo string) (mecideris *Mecideris) {
	mecideris = new(Mecideris)
	mecideris.wheel = wheel
	mecideris.logo = logo
	return
}

//NumberOfWheel will return num of wheels
func (T *Mecideris) NumberOfWheel() int {
	return T.numberOfWheel()
}

//Start will give a message of start
func (T *Mecideris) Start() {
	fmt.Println("I'am :", T.logo)
	T.Car.start()
}

//Stop will give message stop
func (T *Mecideris) Stop() {
	fmt.Println("I'am :", T.logo)
	T.Car.stop()
}

//Cars is type of []*car
type Cars []*Car

//Process process any member of slice with given func
func (cars Cars) Process(f func(car *Car)) {
	for _, c := range cars {
		f(c)
	}
}

//IsFourWheels get if the car have 4 wheels
func IsFourWheels(car *Car) bool {
	if car.numberOfWheel() == 4 {
		return true
	}
	return false
}

//Any is the type of interface
type Any interface{}

// Map will use given func to any menber of the slice
func (cars Cars) Map(f func(car *Car) Any) []Any {
	theresult := make([]Any, 0)
	ix := 0
	cars.Process(func(theCar *Car) {
		theresult[ix] = f(theCar)
		ix++
	})
	return theresult
}

// GetWheels  get the weels of car
func GetWheels(car *Car) Any {
	return car.numberOfWheel()
}

//FindAll return all result suit the conditions
func (cars Cars) FindAll(f func(car *Car) bool) (theCars Cars) {
	theCars = make(Cars, 0)
	cars.Process(func(theCar *Car) {
		if f(theCar) {
			theCars = append(theCars, theCar)
		}
	})
	return
}

//GetNewBMW  will return new BMW
func GetNewBMW(expensiveCars Cars) Cars {
	allNewBMW := expensiveCars.FindAll(func(car *Car) bool {
		return (car.Manufactor == "BMW") && (car.BuildYear > 2010)
	})
	return allNewBMW
}

//ManuSort 采用了闭包,已访问局部切片
func ManuSort(cars Cars) map[string]Cars {
	manuSort := make(map[string]Cars)
	cars.Process(func(car *Car) {
		if _, ok := manuSort[car.Manufactor]; !ok {
			manuSort[car.Manufactor] = make(Cars, 0)
		}
		manuSort[car.Manufactor] = append(manuSort[car.Manufactor], car)
	})
	return manuSort
}

//TestCar is to Test car
func TestCar() {
	ford := &Car{4, "Fiesta", "Ford", 2018, nil}
	ford2 := &Car{4, "X650", "Ford", 2008, nil}
	bmw1 := &Car{4, "S330", "BMW", 2017, nil}
	bmw2 := &Car{4, "Fiesta", "BMW", 2015, nil}
	bmw3 := &Car{4, "Fiesta", "BMW", 2015, nil}
	benz1 := &Car{3, "S650", "Benz", 2018, nil}

	cars := Cars{ford, ford2, bmw1, bmw2, bmw3, benz1}
	// fmt.Println(len(GetNewBMW(cars)))
	mapManu := ManuSort(cars)
	fmt.Println(mapManu)
	fmt.Println("we have", len(mapManu["Benz"]), "BMWs")
}
