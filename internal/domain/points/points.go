package points

type Points struct {
	value int
}

func (s *Points) GetPoints() int {
	return s.value
}

func (s *Points) AddPoints(value int) {
	s.value += value
}
