package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Type  string
	Human // embedded struct
}

func NewHuman(name string, age int) *Human {
	return &Human{
		Name: name,
		Age:  age,
	}
}

func NewAction(actionType string, human *Human) *Action {
	return &Action{
		Type:  actionType,
		Human: *human,
	}
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

func (a *Action) GetType() string {
	return a.Type
}

func main() {
	Rasul := NewHuman("Расул", 23)
	Walk := NewAction("Ходить", Rasul)

	fmt.Println(Walk.GetType())       // Вызов метода Walk
	fmt.Println(Walk.Human.GetName()) // Вызов метода GetName из структуры Human
	fmt.Println(Walk.Human.GetAge())  // Вызов метода GetAge из структуры Human
}
