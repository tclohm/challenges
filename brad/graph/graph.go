package main


type Vertex struct {
	key int
	neighbors map[*Vertex]int
}

func (v *Vertex) AddNeighbor(neighbor Vertex, weight int) {
	v.neighbors[&neighbor] = weight
}

func (v *Vertex) GetConnections() []int {
	keys := []int{}
	for vertex, _ := range v.neighbors {
		keys = append(keys, vertex.key)
	}
	return keys
}

func (v *Vertex) GetWeight(neighbor Vertex) int {
	return v.neighbors[&neighbor]
}

type Graph struct {
	vertices map[int]Vertex
}

func (g *Graph) AddVertex(vertex Vertex) {
	g.vertices[vertex.key] = vertex
}

func (g *Graph) GetVertex(key int) Vertex {
	return g.vertices[key]
}

func (g *Graph) Contains(key int) bool {
	if _, ok := g.vertices[key]; ok {
		return true
	}
	return false
}

func (g *Graph) AddEdge(from, to, weight int) {
	if _, ok := g.vertices[from]; !ok {
		g.AddVertex(Vertex{key: from})
	}
}

func main() {}