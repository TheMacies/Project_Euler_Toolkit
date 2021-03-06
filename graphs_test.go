package kit

import "testing"

func TestDirectedGraph(t *testing.T) {
	g := NewDirectedGraph()

	g.AddVertex()
	g.AddVertex()

	if g.RemoveVertex(-1) == nil || g.RemoveVertex(3) == nil {
		t.Fail()
	}
	if g.AddEdge(0, 2, 1) == nil || g.AddEdge(-2, 1, 3) == nil {
		t.Fail()
	}
	g.AddEdge(0, 1, 5)

	if g.AddEdge(0, 1, 6) == nil {
		t.Fail()
	}

	if g.RemoveVertex(0) != nil {
		t.Fail()
	}

	if len(g.vertices) != 1 {
		t.Fail()
	}

	val, err := g.GetWeight(0, 1)
	if err == nil {
		t.Fail()
	}
	g.AddVertex()
	g.AddEdge(0, 1, 5)
	val, _ = g.GetWeight(0, 1)
	if val != 5 {
		t.Fail()
	}
	g.AddVertex()
	g.AddVertex()
	_, err = g.GetWeight(0, 3)
	if err == nil {
		t.Fail()
	}
	_, err = g.GetWeight(-1, 3)
	if err == nil {
		t.Fail()
	}
	g.AddEdge(2, 3, 5)
	if g.RemoveEdge(1, 4) == nil || g.RemoveEdge(-1, 4) == nil {
		t.Fail()
	}

	if g.AddEdge(1, 1, 1) == nil {
		t.Fail()
	}

	g.RemoveEdge(2, 3)
}

func TestDirectedGraphDijkstra(t *testing.T) {
	g := NewDirectedGraph()
	for i := 0; i < 5; i++ {
		g.AddVertex()
	}

	g.AddEdge(0, 1, 3)
	g.AddEdge(0, 2, 5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 2)
	g.AddEdge(1, 4, 8)
	g.AddEdge(2, 4, 2)
	g.AddEdge(3, 4, 5)

	if l, _ := g.Dijkstra(0); l[1] != 3 {
		t.Fail()
	}
	if l, _ := g.Dijkstra(0); l[4] != 6 {

		t.Fail()
	}

	if l, _ := g.Dijkstra(1); l[2] != 1 {
		t.Fail()
	}

	if _, err := g.Dijkstra(-1); err == nil {
		t.Fail()
	}

	if _, err := g.Dijkstra(5); err == nil {
		t.Fail()
	}

}

func TestUndirectedGraph(t *testing.T) {
	g := NewUndirectedGraph()

	g.AddVertex()
	g.AddVertex()

	if g.RemoveVertex(-1) == nil || g.RemoveVertex(3) == nil {
		t.Fail()
	}
	if g.AddEdge(0, 2, 1) == nil || g.AddEdge(-2, 1, 3) == nil {
		t.Fail()
	}
	g.AddEdge(0, 1, 5)

	if g.AddEdge(0, 1, 6) == nil {
		t.Fail()
	}

	if g.RemoveVertex(0) != nil {
		t.Fail()
	}

	if len(g.vertices) != 1 {
		t.Fail()
	}

	val, err := g.GetWeight(0, 1)
	if err == nil {
		t.Fail()
	}
	g.AddVertex()
	g.AddEdge(0, 1, 5)
	val, _ = g.GetWeight(0, 1)
	if val != 5 {
		t.Fail()
	}
	g.AddVertex()
	g.AddVertex()
	_, err = g.GetWeight(0, 3)
	if err == nil {
		t.Fail()
	}
	_, err = g.GetWeight(-1, 3)
	if err == nil {
		t.Fail()
	}
	g.AddEdge(2, 3, 5)
	if g.RemoveEdge(1, 4) == nil || g.RemoveEdge(-1, 4) == nil {
		t.Fail()
	}
	g.RemoveEdge(2, 3)

	g.AddVertex()
	g.AddVertex()

	g.AddEdge(2, 1, 3)
	g.AddEdge(2, 0, 5)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 5, 3)

	if g.AddEdge(1, 1, 1) == nil {
		t.Fail()
	}
}
