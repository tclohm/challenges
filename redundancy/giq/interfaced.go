package main

import "fmt"

type Person struct {
	Name string
	IsAlive bool
	Children []*Person
}

type Monarchy struct {
	King *Person
	persons map[string]*Person
}

func (this *Monarchy) Init(name string) {
	this.persons = map[string]*Person{
		name: this.King,
	}
}

func (this *Monarchy) Birth(childName, parentName string) {
	var parent = this.persons[parentName]
	var newChild = Person{Name: childName, IsAlive: true}
	parent.Children = append(parent.Children, &newChild)
	this.persons[childName] = &newChild
}

func (this *Monarchy) Death(name string) {
	if person, ok := this.persons[name]; ok {
		person.IsAlive = false
		fmt.Println(name, "has passed away")
	}
}

func (this *Monarchy) getOrderOfSucession() []string {
	var order = []string{}
	var root = this.King
	this.dfs(*root, &order)
	return order
}

func (this *Monarchy) dfs(currentPerson Person, order *[]string) {
	if currentPerson.IsAlive {
		*order = append(*order, currentPerson.Name)
	}
	for i := 0 ; i < len(currentPerson.Children) ; i++ {
		this.dfs(*currentPerson.Children[i], order)
	}
}

func main(){
	m := Monarchy{King: &Person{Name:"A", IsAlive: true, Children: []*Person{}}}
	m.Init("A")

	m.Birth("B", "A")
	m.Birth("C", "A")
	m.Birth("D", "A")
	m.Birth("E", "C")
	m.Birth("F", "C")
	fmt.Println(m.getOrderOfSucession())

	m.Death("A")
	fmt.Println(m.getOrderOfSucession())

	m.Death("E")
	fmt.Println(m.getOrderOfSucession())
}