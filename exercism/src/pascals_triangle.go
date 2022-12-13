package src

func Triangle(n int) [][]int {
	rsp := make([][]int, n)

	for i := 0; i < n; i++ {
		line := make([]int, i+1)

		// First line exception
		if i == 0 {
			line[i] = 1
		}

		// Second lines and below
		for j := 0; j <= i; j++ {
			// Edges of the pyramid are exceptions and always 1
			if j == 0 || j == i {
				line[j] = 1
				continue
			}

			// The inner nodes of the Pyramid are computed as the sum of their predecessors
			line[j] = rsp[i-1][j-1] + rsp[i-1][j]
		}

		rsp[i] = line
	}

	return rsp
}
