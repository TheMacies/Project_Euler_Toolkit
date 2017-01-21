package kit

import (
	"fmt"
	"testing"
)

func TestGetNewPQ(t *testing.T) {
	q := NewPriorityQueue(DefaultIntComparingFuncion)
	if len(q.nodes) != 1 || cap(q.nodes) != 10 {
		fmt.Println(len(q.nodes))
		t.Fail()
	}
}

func TestPQInsert(t *testing.T) {
	q := NewPriorityQueue(DefaultIntComparingFuncion)

	q.Insert(1)
	a, _ := q.nodes[1].(int)
	if a != 1 {
		t.Fail()
	}
	q.Insert(2)
	a, _ = q.nodes[1].(int)
	if a != 2 {
		t.Fail()
	}
	q.Insert(-3)
	a, _ = q.nodes[1].(int)
	if a != 2 {
		fmt.Println(a)
		t.Fail()
	}
}

func TestPQGetPopMax(t *testing.T) {
	q := NewPriorityQueue(DefaultIntComparingFuncion)

	if q.GetMax() != nil {
		t.Fail()
	}

	var val int
	for i := 1; i < 20; i++ {
		q.Insert(i)
	}
	val, _ = q.GetMax().(int)
	if val != 19 {
		t.Fail()
	}

	for i := 19; i >= 10; i-- {
		val, _ = q.PopMax().(int)
		if val != i {
			fmt.Println(val)
			t.Fail()
		}
	}
	for i := 10; i >= 1; i-- {
		q.PopMax()
	}
	if q.PopMax() != nil {
		t.Fail()
	}
}
