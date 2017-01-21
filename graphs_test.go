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
