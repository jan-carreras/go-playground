package union_find

type QuickFind struct {
	id    []int
	count int
	debug Debug
}

type Debugger interface {
	Debug() Debug
}

func NewQuickFind(sites int) *QuickFind {
	id := make([]int, sites)
	for i := range id {
		id[i] = i
	}

	return &QuickFind{
		id:    id,
		count: sites,
	}
}

func (qf *QuickFind) Union(p, q int) {
	pID, qID := qf.Find(p), qf.Find(q)

	if pID == qID {
		return // NoOp if they are connected already
	}

	for i := range qf.id {
		if qf.Find(i) == pID {
			qf.id[i] = qID
		}
	}

	qf.count--
}

func (qf *QuickFind) Find(p int) (componentID int) {
	qf.debug.IDAccesses++
	return qf.id[p]
}

func (qf *QuickFind) Connected(p, q int) (connected bool) {
	return qf.Find(p) == qf.Find(q)
}

func (qf *QuickFind) Count() int {
	return qf.count
}

type Debug struct {
	ID         []int
	IDAccesses int
}

func (qf *QuickFind) Debug() Debug {
	return Debug{
		ID:         qf.id,
		IDAccesses: qf.debug.IDAccesses,
	}
}
