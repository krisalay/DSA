package workAtTech_Array

func setRowColumnZeroes(matrix *[][]int) {
  rowMap, colMap := map[int]bool{}, map[int]bool{}

  for i := 0; i < len(*matrix); i++ {
    for j := 0; j < len(*matrix); j++ {
      if (*matrix)[i][j] == 0 {
        rowMap[i] = true
        colMap[j] = true
      }
    }
  }

  for i := 0; i < len(*matrix); i++ {
    _, rowIsZero := rowMap[i]
    for j := 0; j < len(*matrix); j++ {
      _, colIsZero := rowMap[j]
      if rowIsZero || colIsZero {
        (*matrix)[i][j] = 0
      }
    }
  }
}
