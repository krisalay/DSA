// Source: https://leetcode.com/problems/shortest-path-in-binary-matrix/
package graph

import "fmt"

type Cell struct {
	x, y, distance int
}

type Queue []Cell

func (this *Queue) enqueue(val Cell) {
	*this = append(*this, val)
}

func (this *Queue) isEmpty() bool {
	return len(*this) == 0
}

func (this *Queue) dequeue() (Cell, error) {
	if (*this).isEmpty() {
		return Cell{-1, -1, -1}, fmt.Errorf("queue is empty")
	}
	element := (*this)[0]
	*this = (*this)[1:]
	return element, nil
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	length := len(grid)
	dx := [8]int{-1, 0, 1, -1, 1, -1, 0, 1} // x direction difference
	dy := [8]int{-1, -1, -1, 0, 0, 1, 1, 1} // y direction difference
	firstCell := Cell{x: 0, y: 0, distance: 1}
	queue := &Queue{firstCell}
	visited := make([][]bool, length)
	for i, _ := range visited {
		visited[i] = make([]bool, length)
	}

	visited[0][0] = true

	for !queue.isEmpty() {
		element, _ := queue.dequeue()
		if element.x == length-1 && element.y == length-1 {
			return element.distance
		}
		for i := 0; i < 8; i++ {
			x := element.x + dx[i]
			y := element.y + dy[i]

			if x >= 0 && x < length && y >= 0 && y < length && grid[x][y] == 0 && visited[x][y] == false {
				queue.enqueue(Cell{x: x, y: y, distance: element.distance + 1})
				visited[x][y] = true
			}
		}
	}

	return -1
}

// func main() {
// 	matrix := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
// 	distance := shortestPathBinaryMatrix(matrix)
// 	fmt.Println(distance)
// }
