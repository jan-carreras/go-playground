package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	err := run(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
}

func run(input io.Reader) error {
	forest, err := readForest(input)
	if err != nil {
		return fmt.Errorf("unable to read the forest: %w", err)
	}

	visible := countVisibleTrees(forest)
	fmt.Println("Visible trees:", visible)

	return nil
}

func maxScenicScore(forest [][]int) int {
	maxScore := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest); j++ {
			score := scenicScore(forest, i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

// This doesn't work for trees on the edge of the map. It will always return 0
func scenicScore(forest [][]int, x, y int) int {
	score := 1
	viewDistance := 0
	addToScore := func(distance int) {
		if distance != 0 {
			score *= distance
		}
		viewDistance = 0 // Resets the viewDistance
	}

	// top
	for i := x - 1; i >= 0; i-- {
		//fmt.Println("up", forest[i][y])
		viewDistance++
		if forest[i][y] >= forest[x][y] {
			break
		}
	}
	addToScore(viewDistance)

	// down
	for i := x + 1; i < len(forest); i++ {
		//fmt.Println("down", forest[i][y])
		viewDistance++
		if forest[i][y] >= forest[x][y] {
			break
		}
	}
	addToScore(viewDistance)

	// left
	for i := y - 1; i >= 0; i-- {
		//fmt.Println("left", forest[x][i])
		viewDistance++
		if forest[x][i] >= forest[x][y] {
			break
		}
	}
	addToScore(viewDistance)

	// right
	for i := y + 1; i < len(forest); i++ {
		//fmt.Println("right", forest[x][i])
		viewDistance++
		if forest[x][i] >= forest[x][y] {
			break
		}
	}
	addToScore(viewDistance)

	return score
}

func countVisibleTrees(forest [][]int) int {
	return visibleTrees(
		seenFromLeft(forest),
		seenFromRight(forest),
		seenFromDownTop(forest),
		seenTopDown(forest),
	)
}

func visibleTrees(views ...[][]int) int {
	notVisible := 0
	for i := 0; i < len(views[0]); i++ {
		for j := 0; j < len(views[0]); j++ {
			rowNotVisible := 0
			for x := 0; x < len(views); x++ {
				if views[x][i][j] == -1 {
					rowNotVisible++
				}
			}
			notVisible += rowNotVisible / 4
		}
	}

	return (len(views[0]) * len(views[0])) - notVisible
}

func seenFromDownTop(forest [][]int) [][]int {
	down := copyForest(forest)
	for i := 0; i < len(down); i++ {
		max := 0
		for j := len(down) - 1; j > 0; j-- {
			if down[j][i] <= max {
				if j != 0 && j != len(down)-1 {
					down[j][i] = -1
				}
			} else {
				max = down[j][i]
			}
		}
	}

	return down
}

func seenTopDown(forest [][]int) [][]int {
	top := copyForest(forest)
	for i := 0; i < len(top); i++ {
		max := 0
		for j := 0; j < len(top); j++ {
			if top[j][i] <= max {
				if j != 0 && j != len(top)-1 {
					top[j][i] = -1
				}
			} else {
				max = top[j][i]
			}
		}
	}
	return top
}

func seenFromRight(forest [][]int) [][]int {
	right := copyForest(forest)
	for i := 0; i < len(right); i++ {
		max := 0
		for j := len(right) - 1; j > 0; j-- {
			if right[i][j] <= max {
				if j != 0 && j != len(right)-1 {
					right[i][j] = -1
				}
			} else {
				max = right[i][j]
			}
		}
	}
	return right

}

func seenFromLeft(forest [][]int) [][]int {
	left := copyForest(forest)
	for i := 0; i < len(left); i++ {
		max := 0
		for j := 0; j < len(left); j++ {
			if left[i][j] <= max {
				if j != 0 && j != len(left)-1 {
					left[i][j] = -1
				}
			} else {
				max = left[i][j]
			}
		}
	}
	return left
}

func readForest(input io.Reader) ([][]int, error) {
	buf := bufio.NewScanner(input)
	forest := make([][]int, 0)
	for buf.Scan() {
		t := buf.Text()
		lineOfTrees := make([]int, 0, len(t))
		for _, c := range t {
			treeHight, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, fmt.Errorf("invalid tree high: %q", c)
			}
			lineOfTrees = append(lineOfTrees, treeHight)
		}

		forest = append(forest, lineOfTrees)
	}

	return forest, buf.Err()
}

func copyForest(forest [][]int) [][]int {
	newForest := make([][]int, 0, len(forest))
	for i := 0; i < len(forest); i++ {
		line := make([]int, len(forest[i]))
		copy(line, forest[i])
		newForest = append(newForest, line)
	}
	return newForest
}

func printForest(forest [][]int) {
	for _, line := range forest {
		fmt.Println(line)
	}
	fmt.Println()
}
