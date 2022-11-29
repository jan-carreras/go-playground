package union_find

// 1.5.12 Quick-union with path compression. Modify quick-union (page 224) to
// include path compression, by adding a loop to union() that links every site on
// the paths from p and q to the roots of their trees to the root of the new
// tree. Give a sequence of input pairs that causes this method to produce a path
// of length 4. Note : The amortized cost per operation for this algorithm is
// known to be logarithmic.

type WeightedQuickUnionPathCompression struct {
	ids   []int
	size  []int
	count int
}

func NewWeightedQuickUnionPathCompression(sites int) *WeightedQuickUnionPathCompression {
	ids := make([]int, sites)
	for i := range ids {
		ids[i] = i
	}

	size := make([]int, sites)
	for i := range size {
		size[i] = 1
	}

	return &WeightedQuickUnionPathCompression{
		ids:   ids,
		size:  size,
		count: sites,
	}
}

func (w *WeightedQuickUnionPathCompression) Union(p, q int) {
	i := w.Find(p)
	j := w.Find(q)
	if i == j {
		return
	}

	if w.size[i] < w.size[j] {
		w.pathCompression(p, q)
		w.ids[i] = j
		w.size[j] += w.size[i]
	} else {
		w.pathCompression(q, p)
		w.ids[j] = i
		w.size[i] += w.size[j]
	}

	w.count--
}

func (w *WeightedQuickUnionPathCompression) pathCompression(p, q int) {
	rootP, rootQ := w.Find(p), w.Find(q)
	// I want to move up the three of P and assign every note and point to rootQ
	for p != rootP {
		w.ids[p], p = rootQ, w.ids[p]
	}

	w.ids[rootP] = q
}

func (w *WeightedQuickUnionPathCompression) Find(p int) (componentID int) {
	for p != w.ids[p] {
		p = w.ids[p]
	}

	return p
}

func (w *WeightedQuickUnionPathCompression) Connected(p, q int) (connected bool) {
	return w.Find(p) == w.Find(q)
}

func (w *WeightedQuickUnionPathCompression) Count() int {
	return w.count
}
