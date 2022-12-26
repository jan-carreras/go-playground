package dynamic_programming

import "fmt"

type Item struct {
	name   string
	weight int
	score  int
}

func (i Item) symbol() string {
	if len(i.name) == 0 {
		return "?"
	}
	return string(i.name[0])
}

type point struct {
	score int
	name  string
}

func newPoint(name string, score int) point {
	return point{
		name:  name,
		score: score,
	}
}

func itemsToTake(items []Item, maxWeight int) point {
	table := makeEmptyTable(items, maxWeight+1) // Don't use the index 0; makes array arithmetic easier

	for i, item := range items {
		for w := 1; w <= maxWeight; w++ {
			if i == 0 { // First row is special since doesn't have previous row
				if item.weight <= w {
					table[i][w] = newPoint(item.symbol(), item.score)
				}
				continue
			}

			var combined point
			if w-item.weight >= 0 {
				a := table[i-1][w-item.weight]
				combined = newPoint(a.name+item.symbol(), a.score+item.score)
			}

			previous := table[i-1][w]
			table[i][w] = max(previous, combined)
		}
	}

	printTable(table)
	return table[len(items)-1][maxWeight-1]
}

func max(a, b point) point {
	if a.score > b.score {
		return a
	}
	return b
}

func makeEmptyTable(items []Item, maxWeight int) [][]point {
	table := make([][]point, len(items))
	for i := range table {
		table[i] = make([]point, maxWeight)
	}
	return table
}

func printTable(t [][]point) {
	for i := range t {
		for j := range t[i] {
			if j == 0 {
				continue
			}
			fmt.Printf("%2.d(%3s)\t", t[i][j].score, t[i][j].name)
		}
		fmt.Println()
	}
}
