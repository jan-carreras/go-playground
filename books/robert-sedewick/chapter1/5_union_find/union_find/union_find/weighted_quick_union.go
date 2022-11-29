package union_find

type WeightedQuickUnion struct {
	ids   []int
	size  []int
	count int
	debug Debug
}

func NewWeightedQuickUnion(sites int) *WeightedQuickUnion {
	ids := make([]int, sites)
	for i := range ids {
		ids[i] = i
	}

	size := make([]int, sites)
	for i := range size {
		size[i] = 1
	}

	return &WeightedQuickUnion{
		ids:   ids,
		size:  size,
		count: sites,
	}
}

func (w *WeightedQuickUnion) Union(p, q int) {
	i := w.Find(p)
	j := w.Find(q)
	if i == j {
		return
	}

	if w.size[i] < w.size[j] {
		w.ids[i] = j
		w.size[j] += w.size[i]
	} else {
		w.ids[j] = i
		w.size[i] += w.size[j]
	}
	w.debug.IDAccesses += 3

	w.count--
}

func (w *WeightedQuickUnion) Find(p int) (componentID int) {
	for p != w.ids[p] {
		w.debug.IDAccesses += 2
		p = w.ids[p]
	}

	return p
}

func (w *WeightedQuickUnion) Connected(p, q int) (connected bool) {
	return w.Find(p) == w.Find(q)
}

func (w *WeightedQuickUnion) Count() int {
	return w.count
}

func (w *WeightedQuickUnion) Debug() Debug {
	d := w.debug
	d.ID = w.ids
	return d
}
