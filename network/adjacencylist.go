package main

import "fmt"

type SocialGraph struct {
	Size 	int
	Links 	[][]Link
}

type Link struct {
	Vertex1 	int
	Vertex2 	int 
	LinkWeight 	 int
}

func NewSocialGraph(num int) *SocialGraph {
	reutrn &SocialGraph{
		Size: num,
		Links: make([][]Link, num),
	}
}

func (graph *SocialGraph) AddLink(v1 int, v2 int, weight int) {
	graph.Links[v1] = append(graph.Links[v1], Link{Vertex1: v1, Vertex2: v2, LinkWeight: weight})
} 

func (graph *SocialGraph) PrintLinks() {
	var vertex int = 0
	fmt.Println("Printing all links from %d\n", vertex)

	for _, link := range graph.Links[vertex] {
		for
	}
}