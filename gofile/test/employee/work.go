package employee
import (
	"time"
)
type Employee struct{
	birthday int
	salary float32
	name string
	grade int8
}

func (T *Employee)GiveRaise(rate float32,grade int8){
	T.salary *= rate
	T.grade += grade
}

func (T *Employee)Salary()(float32){
	return T.salary
}

func (T *Employee)SetSalary(salary float32){
	T.salary = salary
}

func (T * Employee)Grade()int8{
	return T.grade
}

func (T * Employee)Age()int{
	return time.Now().Year()-T.birthday
}

func NewEmployee(salary float32,name string, birthday int, grade int8)(*Employee){
	employee := new(Employee)
	employee.birthday = birthday
	employee.name = name
	employee.salary = salary
	employee.grade = grade
	return employee
}
