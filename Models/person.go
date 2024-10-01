package Models

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPerson(id int, name string, age int) *Person {
	return &Person{
		Id:   id,
		Name: name,
		Age:  age,
	}
}
