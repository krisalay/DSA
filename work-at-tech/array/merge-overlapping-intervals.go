package main

import "fmt"

func mergeIntervals(intervals [][]int) {
  // Get Maximum Value
  maxPoint := intervals[0][0]
  for _, val := range intervals {
    endPoint := val[1]
    maxPoint = max(endPoint, maxPoint)
  }

  // Create a zero filled array with length of maxPoint + 1
  track := make([]int, maxPoint+1)

  // Populate the track array with ones
  for _, val := range intervals {
    for i := val[0]; i < val[1]; i++ {
      track[i] = 1
    }
  }

  // Iterate the track array and find the overlapping
  result := [][]int{}
  temp := []int{}
  for i, val := range track {
    if val == 0 {
      if len(temp) != 0 {
        result = append(result, temp)
        temp = []int{}
      }
    } else {
      if len(temp) == 0 {
        temp = append(temp, i)
        temp = append(temp, i + 1)
      } else {
        temp[1] = i+1
      }
    }
  }


  fmt.Println(result)

}

func max(a, b int) int {
  if a >= b {
    return a
  }
  return b
}

func main() {
  mergeIntervals([][]int{ {1, 2}, {2, 3}, {1, 4}, {5, 6} })
  mergeIntervals([][]int{ {1, 4}, {6, 7}, {4, 5} })
}
