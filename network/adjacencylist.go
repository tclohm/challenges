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
	return &SocialGraph{
		Size: num,
		Links: make([][]Link, num),
	}
}

func (graph *SocialGraph) AddLink(v1 int, v2 int, weight int) {
	graph.Links[v1] = append(graph.Links[v1], Link{Vertex1: v1, Vertex2: v2, LinkWeight: weight})
} 

func (graph *SocialGraph) PrintLinks() {
	var vertex int = 0
	fmt.Printf("Printing all links from %d\n", vertex)

	for _, link := range graph.Links[vertex] {
		fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
	}

	fmt.Println("Printing all links in graph")

	for _, adjacent := range graph.Links {
		for _, link := range adjacent {
			fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
		}
	}
}

func main() {
	var social *SocialGraph

	social = NewSocialGraph(4)

	social.AddLink(0, 1, 1)
	social.AddLink(0, 2, 1)
	social.AddLink(1, 3, 1)
	social.AddLink(2, 4, 1)

	social.PrintLinks()

}