package kit

import (
	"fmt"
	"testing"
)

func TestInitializeSet(t *testing.T) {
	InitializeSet(DefaultIntComparingFuncion)
}

func TestDefaultByteCompaingFunction(t *testing.T) {
	if DefaultByteComparingFuncion(byte(2), byte(2)) != 0 {
		t.Fail()
	}

	if DefaultByteComparingFuncion(byte(1), byte(5)) != -1 {
		t.Fail()
	}

	if DefaultByteComparingFuncion(byte(50), byte(8)) != 1 {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	s := InitializeSet(DefaultByteComparingFuncion)
	if s.GetMax() != nil {
		t.Fail()
	}
	s.Insert(byte(2))
	if s.GetMax() != byte(2) {
		t.Fail()
	}
	s.Insert(byte(5))
	if s.GetMax() != byte(5) {
		fmt.Println(s.GetMax())
		t.Fail()
	}
	s.Insert(byte(0))
	if s.GetMax() != byte(5) {
		t.Fail()
	}
	size := len(s.elements)
	s.Insert(byte(2))
	if size != len(s.elements) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	s := InitializeSet(DefaultIntComparingFuncion)
	s.Insert(2)
	s.Delete(2)
	if len(s.elements) != 0 {
		t.Fail()
	}
}

func TestGetMinMax(t *testing.T) {
	s := InitializeSet(DefaultStringComparingFuncion)
	s.Insert("ab")
	s.Insert("ab")
	s.Insert("a")
	s.Insert("zz")
	if s.GetMax() != "zz" {
		t.Fail()
	}
	if s.GetMin() != "a" {
		t.Fail()
	}
}

func TestPopMinMax(t *testing.T) {
	s := InitializeSet(DefaultIntComparingFuncion)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.PopMax()
	s.PopMin()
	if s.GetMax() != s.GetMin() {
		t.Fail()
	}
	s.PopMax()
	if s.PopMax() != nil {
		t.Fail()
	}
	if s.PopMin() != nil {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	s := InitializeSet(DefaultIntComparingFuncion)
	s.Insert(2)
	s.Insert(4)
	if !s.Contains(2) {
		t.Fail()
	}
	if !s.Contains(4) {
		t.Fail()
	}
	if s.Contains(3) {
		t.Fail()
	}
}

func TestGetElems(t *testing.T) {
	s := InitializeSet(DefaultIntComparingFuncion)
	s.Insert(2)
	s.Insert(4)
	a := s.GetElems()
	val := a[0].(int)
	if val != 2 {
		t.Fail()
	}
	val = a[1].(int)
	if val != 4 {
		t.Fail()
	}
}
