package core

type State[K string | int, V any] struct {
	States map[K]V
}

func (s *State[K, V]) Set(k K, v V) {
	if s.States == nil {
		s.States = map[K]V{}
	}

	s.States[k] = v
}

func (s *State[K, V]) Get(k K) V {
	return s.States[k]
}
