package kit

type Set struct {
	elements []interface{}
	// The comparing function shall return 0 if a == b equal  > 0  if a > b and  < 0 if a < b
	comparingFunction func(interface{}, interface{}) int
}

func (s *Set) GetElems() []interface{} {
	return s.elements
}

func InitializeSet(f func(interface{}, interface{}) int) Set {
	var s Set
	s.comparingFunction = f
	return s
}

func (s *Set) Insert(el interface{}) {
	for i, val := range s.elements {
		eval := s.comparingFunction(el, val)
		if eval == 0 {
			return
		}
		if eval < 0 {
			s.elements = append(s.elements[:i], append([]interface{}{el}, s.elements[i:]...)...)
			return
		}
	}
	s.elements = append(s.elements, el)
}

func (s *Set) Delete(el interface{}) {
	for i, val := range s.elements {
		eval := s.comparingFunction(el, val)
		if eval == 0 {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			return
		}
	}
}

func (s *Set) GetMax() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

func (s *Set) PopMax() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	el := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-2]
	return el
}

func (s *Set) GetMin() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[0]
}

func (s *Set) PopMin() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	el := s.elements[0]
	s.elements = s.elements[1:]
	return el
}

func (s *Set) Contains(el interface{}) bool {
	for _, val := range s.elements {
		eval := s.comparingFunction(el, val)
		if eval == 0 {
			return true
		}
	}
	return false
}

func DefaultIntComparingFuncion(a, b interface{}) int {
	aInt := a.(int)
	bInt := b.(int)
	c := bInt - aInt
	if c == 0 {
		return c
	}
	if c > 0 {
		return 1
	}
	return -1
}

func DefaultStringComparingFuncion(a, b interface{}) int {
	aString := a.(string)
	bString := b.(string)

	if aString > bString {
		return 1
	}
	if aString == bString {
		return 0
	}
	return -1
}

func DefaultByteComparingFuncion(a, b interface{}) int {
	aByte := a.(byte)
	bByte := b.(byte)

	if aByte > bByte {
		return 1
	}
	if aByte == bByte {
		return 0
	}
	return -1
}
