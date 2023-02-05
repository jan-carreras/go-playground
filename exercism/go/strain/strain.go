package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (rsp Ints) {
	return keep[int](i, filter)
}

func (i Ints) Discard(filter func(int) bool) Ints {
	reverseFnx := func(i int) bool { return !filter(i) }
	return i.Keep(reverseFnx)
}

func (l Lists) Keep(filter func([]int) bool) (rsp Lists) {
	return keep[[]int](l, filter)
}

func (s Strings) Keep(filter func(string) bool) (rsp Strings) {
	return keep[string](s, filter)
}

func keep[T any](list []T, filter func(T) bool) (rsp []T) {
	for _, value := range list {
		if filter(value) {
			rsp = append(rsp, value)
		}
	}
	return rsp
}
