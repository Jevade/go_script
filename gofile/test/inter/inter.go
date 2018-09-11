package inter
type Namer interface{
	Get()string
	Set(string)
}

type Person struct{
    name string
	sex bool
	age int8
	length float32
	weight float32
	nationality string
	position string
}
func(T *Person)Get()string{
    return T.name
}
func(T *Person)Set(name string){
    T.name = name
}
func (T *Person)String()string{
    return T.name + ";" + T.position
}
func  NewPerson(name,position string)(T *Person){
    person := new(Person)
	person.name = name
	person.position = position
	return person
}
type Cat struct{
    name string
	age int
}

func (T Cat)Get()string{
    return T.name
}
func (T *Cat)Set(name string){
    T.name = name
}



