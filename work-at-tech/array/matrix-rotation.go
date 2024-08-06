package workAtTech_Array

func rotateMatrix(matrix [][]int) [][]int {
  // Setup resulting matrix
  resultRow, resultCol := len(matrix[0]), len(matrix)
  result := make([][]int, resultRow)
  for i, _ := range result {
    result[i] = make([]int, resultCol)
  }

  // Rotate Matrix
  for j := 0; j < len(matrix[0]); j++ {
    for i := len(matrix) - 1; i >= 0; i-- {
      result[j][len(matrix) - 1 - i] = matrix[i][j]
    }
  }

  return result
}
