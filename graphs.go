package kit

import (
	"errors"
	"math"
)

type edge struct {
	weight         int
	adjacentVertex int
}

type vertex struct {
	edges []edge
}

type graph struct {
	vertices []vertex
}

type undirectedGraph struct {
	graph
}

type directedGraph struct {
	graph
}

//undirectedGraph implementation
func NewUndirectedGraph() *undirectedGraph {
	return &undirectedGraph{}
}

func (g *undirectedGraph) AddVertex() int {
	g.vertices = append(g.vertices, vertex{})
	return len(g.vertices) - 1
}

func (g *undirectedGraph) RemoveVertex(number int) error {
	if number < 0 {
		return errors.New("Bad vertex number ( numbers are >= 0 )")
	}
	if number >= len(g.vertices) {
		return errors.New("Vertex does not exist")
	}
	for _, edge := range g.vertices[number].edges {
		g.RemoveEdge(number, edge.adjacentVertex)
		g.RemoveEdge(edge.adjacentVertex, number)
	}

	g.vertices = append(g.vertices[:number], g.vertices[number+1:]...)
	return nil
}

func (g *undirectedGraph) AddEdge(vertex1, vertex2, weight int) error {
	if vertex1 < 0 || vertex2 < 0 {
		return errors.New("Bad vertex number ( numbers are >= 0 )")
	}
	if l := len(g.vertices); l <= vertex1 || l <= vertex2 {
		return errors.New("Vertex does not exist")
	}
	if vertex1 == vertex2 {
		return errors.New("No loops allowed")
	}
	if len(g.vertices[vertex1].edges) > len(g.vertices[vertex2].edges) {
		vertex1, vertex2 = vertex2, vertex1
	}
	for _, edge := range g.vertices[vertex1].edges {
		if edge.adjacentVertex == vertex2 {
			return errors.New("Edge exists")
		}
	}
	g.vertices[vertex1].edges = append(g.vertices[vertex1].edges,
		edge{adjacentVertex: vertex2, weight: weight})
	g.vertices[vertex2].edges = append(g.vertices[vertex2].edges,
		edge{adjacentVertex: vertex1, weight: weight})
	return nil
}

func (g *undirectedGraph) RemoveEdge(vertex1, vertex2 int) error {
	if vertex1 < 0 || vertex2 < 0 {
		return errors.New("Bad vertex number ( numbers are >= 0 )")
	}
	if l := len(g.vertices); l <= vertex1 || l <= vertex2 {
		return errors.New("Vertex does not exist")
	}

	for i, edge := range g.vertices[vertex1].edges {
		if edge.adjacentVertex == vertex2 {
			g.vertices[vertex1].edges = append(g.vertices[vertex1].edges[:i], g.vertices[vertex1].edges[i+1:]...)
			for j, edge := range g.vertices[vertex2].edges {
				if edge.adjacentVertex == vertex1 {
					g.vertices[vertex2].edges = append(g.vertices[vertex2].edges[:j], g.vertices[vertex2].edges[j+1:]...)
					break
				}
			}
			break
		}
	}

	return nil
}

func (g *undirectedGraph) GetWeight(vertex1, vertex2 int) (int, error) {
	if vertex1 < 0 || vertex2 < 0 {
		return 0, errors.New("Bad vertex number ( numbers are >= 0 )")
	}

	if l := len(g.vertices); l <= vertex1 || l <= vertex2 {
		return 0, errors.New("Vertex does not exist")
	}

	for _, edge := range g.vertices[vertex1].edges {
		if edge.adjacentVertex == vertex2 {
			return edge.weight, nil
		}
	}
	return 0, errors.New("Edge does not exist")
}

//directedGraph implementation
func NewDirectedGraph() *directedGraph {
	return &directedGraph{}
}

func (g *directedGraph) AddVertex() int {
	g.vertices = append(g.vertices, vertex{})
	return len(g.vertices) - 1
}

func (g *directedGraph) RemoveVertex(number int) error {
	if number < 0 {
		return errors.New("Bad vertex number ( numbers are >= 0 )")
	}
	if number >= len(g.vertices) {
		return errors.New("Vertex does not exist")
	}

	for i := range g.vertices {
		g.RemoveEdge(i, number)
	}

	g.vertices = append(g.vertices[:number], g.vertices[number+1:]...)
	return nil
}

func (g *directedGraph) AddEdge(vertex1, vertex2, weight int) error {
	if vertex1 < 0 || vertex2 < 0 {
		return errors.New("Bad vertex number ( numbers are >= 0 )")
	}
	if l := len(g.vertices); l <= vertex1 || l <= vertex2 {
		return errors.New("Vertex does not exist")
	}

	if vertex1 == vertex2 {
		return errors.New("No loops allowed")
	}

	for _, edge := range g.vertices[vertex1].edges {
		if edge.adjacentVertex == vertex2 {
			return errors.New("Edge exists")
		}
	}

	g.vertices[vertex1].edges = append(g.vertices[vertex1].edges,
		edge{adjacentVertex: vertex2, weight: weight})

	return nil
}

func (g *directedGraph) RemoveEdge(vertex1, vertex2 int) error {
	if vertex1 < 0 || vertex2 < 0 {
		return errors.New("Bad vertex number ( numbers are >= 0 )")
	}

	if l := len(g.vertices); l <= vertex1 || l <= vertex2 {
		return errors.New("Vertex does not exist")
	}

	for i, edge := range g.vertices[vertex1].edges {
		if edge.adjacentVertex == vertex2 {
			g.vertices[vertex1].edges = append(g.vertices[vertex1].edges[:i], g.vertices[vertex1].edges[i+1:]...)
			break
		}
	}
	return nil
}

func (g *directedGraph) GetWeight(vertex1, vertex2 int) (int, error) {
	if vertex1 < 0 || vertex2 < 0 {
		return 0, errors.New("Bad vertex number ( numbers are >= 0 )")
	}

	if l := len(g.vertices); l <= vertex1 || l <= vertex2 {
		return 0, errors.New("Vertex does not exist")
	}

	for _, edge := range g.vertices[vertex1].edges {
		if edge.adjacentVertex == vertex2 {
			return edge.weight, nil
		}
	}
	return 0, errors.New("Edge does not exist")
}

func (g *directedGraph) Dijkstra(vertex1 int) ([]int, error) {
	type visited struct {
		dist    int
		visited bool
	}

	if vertex1 < 0 {
		return nil, errors.New("Bad vertex number ( numbers are >= 0 )")
	}

	if l := len(g.vertices); l <= vertex1 {
		return nil, errors.New("Vertex does not exist")
	}

	v := make([]visited, len(g.vertices))
	verticesLeft := len(g.vertices)

	for i := range g.vertices {
		if i == vertex1 {
			v[i] = visited{0, false}
		} else {
			v[i] = visited{math.MaxInt32, false}
		}
	}

	for verticesLeft > 0 {
		var indClosest int
		closeness := math.MaxInt32
		for i := 0; i < len(g.vertices); i++ {
			if !v[i].visited && v[i].dist <= closeness {
				closeness = v[i].dist
				indClosest = i
			}
		}
		v[indClosest].visited = true
		verticesLeft--
		for _, edge := range g.vertices[indClosest].edges {
			if alt := edge.weight + v[indClosest].dist; !v[edge.adjacentVertex].visited && alt < v[edge.adjacentVertex].dist {
				v[edge.adjacentVertex].dist = alt
			}
		}
	}
	lengths := make([]int, len(g.vertices))
	for i := range v {
		lengths[i] = v[i].dist
	}
	return lengths, nil
}
