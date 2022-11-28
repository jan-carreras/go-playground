package union_find

type QuickUnion struct {
	ids   []int
	count int
}

func NewQuickUnion(sites int) *QuickUnion {
	ids := make([]int, sites)
	for i := range ids {
		ids[i] = i
	}

	return &QuickUnion{
		ids:   ids,
		count: sites,
	}
}

func (q2 *QuickUnion) Union(p, q int) {
	pRoot, qRoot := q2.Find(p), q2.Find(q)
	if pRoot == qRoot {
		return
	}

	q2.ids[pRoot] = qRoot
	q2.count--
}

func (q2 *QuickUnion) Find(p int) (componentID int) {
	for q2.ids[p] != p {
		p = q2.ids[p]
	}

	return p
}

func (q2 *QuickUnion) Connected(p, q int) (connected bool) {
	return q2.Find(p) == q2.Find(q)
}

func (q2 *QuickUnion) Count() int {
	return q2.count
}
