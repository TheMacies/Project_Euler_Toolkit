package kit

import "testing"

func TestGetMax(t *testing.T) {
	var q PriorityQueue
	if _, err := q.GetMax(); err == nil {
		t.Fail()
	}
	q.Insert(1)
	if val, _ := q.GetMax(); val != 1 {
		t.Fail()
	}
	q.Insert(10)
	if val, _ := q.GetMax(); val != 10 {
		t.Fail()
	}
	q.Insert(5)
	if val, _ := q.GetMax(); val != 10 {
		t.Fail()
	}
	q.Insert(4)
	if val, _ := q.GetMax(); val != 10 {
		t.Fail()
	}
}
