package person

// Person represents the person type
type Person struct {
	name  string
	age   uint
	Hobby string
}

func NewPerson(name string, age uint, hobby string) Person {
	return Person{name: name, age: age, Hobby: hobby}
}
func (p Person) Name() string {
	return p.name
}
func (p Person) Age() uint {
	return p.age
}
