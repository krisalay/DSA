package DP

// import "fmt"

func pascalsTriangle(numRows int) [][]int {
  // Create a 2D DP array
  dp := make([][]int, numRows+1)
  for i, _ := range dp {
    dp[i] = make([]int, numRows+1)
  }

  dp[1][1] = 1
  result := [][]int{{1}}

  for i := 2; i <= numRows; i++ {
    temp := []int{}
    for j := 0; j < i; j++ {
      dp[i-j][j+1] = dp[i-j][j] + dp[i-j-1][j+1]
      temp = append(temp, dp[i-j][j+1])
    }
    result = append(result, temp)
  }
  return result
}

// func main() {
//   result := pascalsTriangle(7)
//   fmt.Println(result)
// }
