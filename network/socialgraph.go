package main

import "fmt"

type Name string

type SocialGraph struct {
	GraphNodes map[Name]struct{}
	Links map[Name]map[Name]struct{}
}

func NewSocialGraph() *SocialGraph {
	return &SocialGraph{
		GraphNodes: make(map[Name]struct{}),
		Links: make(map[Name]map[Name]struct{}),
	}
}

func (graph *SocialGraph) AddEntity(name Name) bool {
	if _, exists := graph.GraphNodes[name]; exists {
		return true
	}

	graph.GraphNodes[name] = struct{}{}
	return true
}

func (graph *SocialGraph) AddLink(name1 Name, name2 Name) {
	if _, exists := graph.GraphNodes[name1]; !exists {
		graph.AddEntity(name1)
	}
	if _, exists := graph.GraphNodes[name2]; !exists {
		graph.AddEntity(name2)
	}
	if _, exists := graph.Links[name1]; !exists {
		graph.Links[name1] = make(map[Name]struct{})
	}
	graph.Links[name1][name2] = struct{}{}
}

func (graph *SocialGraph) PrintLinks() {
	root := Name("Root")

	fmt.Printf("Printing all links adjacent to %d\n", root)

	for node := range graph.Links[root] {
		fmt.Printf("Link: %d -> %d\n", root, node)
	}

	fmt.Println("Printing all links.")

	for r, m := range graph.Links {
		for vertex := range m {
			fmt.Printf("Links: %d -> %d\n", r, vertex)
		}
	}
}

func main() {
	var graph *SocialGraph

	graph = NewSocialGraph()

	var root Name = Name("Root")
	var jon Name = Name("Jon Smith")
	var per Name = Name("Per Jambeck")
	var cynthia Name = Name("Cynthia Gibs")

	graph.AddEntity(root)
	graph.AddEntity(jon)
	graph.AddEntity(per)
	graph.AddEntity(cynthia)

	
	graph.AddLink(root, jon)
	graph.AddLink(root, per)
	graph.AddLink(root, cynthia)

	var mayo Name = Name("Mayo Smith")
	var lorrie Name = Name("Lorrie Jambeck")
	var ellie Name = Name("Ellie Block")

	graph.AddLink(jon, mayo)
	graph.AddLink(jon, lorrie)
	graph.AddLink(per, ellie)

	graph.PrintLinks()
}