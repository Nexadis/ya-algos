package set

type mySet struct {
	data [][]int
	size int
}

func NewSet(size int) mySet {
	set := mySet{
		size: size,
	}
	data := make([][]int, size)
	for i := range data {
		data[i] = make([]int, 0, 10)
	}
	set.data = data
	return set
}

func (s *mySet) Add(x int) {
	if s.Find(x) {
		return
	}
	s.data[x%s.size] = append(s.data[x%s.size], x)
}

func (s *mySet) Find(x int) bool {
	l := s.data[x%s.size]
	for _, v := range l {
		if x == v {
			return true
		}
	}
	return false
}

func (s *mySet) Delete(x int) {
	l := s.data[x%s.size]
	for i, v := range l {
		if x == v {
			l[i] = l[len(l)-1]
			l = l[:len(l)-1]
		}
	}
}
