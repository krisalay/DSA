package graph

import (
// 	"fmt"
	"sort"
)

func makeSet(N int, parent *[]int) {
	for i := 1; i <= N; i++ {
		(*parent)[i] = i
	}
}

func find(i int, parent *[]int) int {
	for (*parent)[i] != i {
		i = (*parent)[i]
	}
	return i
}

type SortByWeight [][]int

func (this SortByWeight) Len() int           { return len(this) }
func (this SortByWeight) Less(i, j int) bool { return this[i][2] < this[j][2] }
func (this SortByWeight) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

func MinimumSpanningTreeCost(N int, edges [][]int) int {
	// Make set
	parent := make([]int, N+1)
	makeSet(N, &parent)
	cost := 0
	// Sort Edge
	sort.Sort(SortByWeight(edges))
	for _, edge := range edges {
		source := edge[0]
		destination := edge[1]
		sourceRepresentative := find(source, &parent)
		destinationRepresentative := find(destination, &parent)
		if sourceRepresentative != destinationRepresentative { // No cycle
			cost += edge[2]
			parent[sourceRepresentative] = destinationRepresentative // Union
		}
	}

	return cost
}

// func main() {
// 	cost := MinimumSpanningTreeCost(4, [][]int{{1, 2, 1}, {2, 3, 4}, {1, 4, 3}, {4, 3, 2}, {1, 3, 10}})
// 	fmt.Println(cost)
// }
