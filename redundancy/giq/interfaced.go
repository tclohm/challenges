package main

import "fmt"

type Person struct {
	Name string
	IsAlive bool
	Children []*Person
}

type Monarchy struct {
	King Person
	persons map[string]Person
}

func (this *Monarchy) Init() {
	this.persons = map[string]Person{
		"king": this.King,
	}
}

func (this *Monarchy) Birth(childName, parentName string) {
	var parent = this.persons[parentName]
	var newChild = Person{Name: childName, IsAlive: true}
	parent.Children = append(parent.Children, &newChild)
	this.persons[childName] = newChild
}

func (this *Monarchy) Death(name string) {
	if person, ok := this.persons[name]; ok {
		person.IsAlive = false
	}
}

func (this *Monarchy) getOrderOfSucession() {
	var order = []string{}
	this.dfs(this.King, &order)
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
	m := Monarchy{King: Person{Name:"M", IsAlive: true, Children: []*Person{}}}
	m.Init()
	fmt.Println(m)

	m.Birth("Henry", "M")

	fmt.Println(m)
}