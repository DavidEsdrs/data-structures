package sets

type Set[Key comparable, Value any] struct {
	set map[Key]Value
}

func NewSet[Key comparable, Value any]() Set[Key, Value] {
	return Set[Key, Value]{
		set: make(map[Key]Value),
	}
}

func (s *Set[Key, Value]) Add(key Key, value Value) {
	if _, exists := s.set[key]; !exists {
		s.set[key] = value
	}
}

func (s *Set[Key, Value]) Get(key Key) (value Value, exists bool) {
	value, exists = s.set[key]
	return value, exists
}

func (s *Set[Key, Value]) Has(key Key) bool {
	_, exists := s.set[key]
	return exists
}

func (s *Set[Key, Value]) Delete(key Key) {
	delete(s.set, key)
}
