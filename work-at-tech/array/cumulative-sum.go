package workAtTech_Array

func GetCumulativeSum(arr []int) []int {
  result := make([]int, len(arr))
  tempSum := 0
  for i, num := range arr {
    tempSum+=num
    result[i] = tempSum
  }
  return result
}
