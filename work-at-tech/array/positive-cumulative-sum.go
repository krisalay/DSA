package workAtTech_Array

func GetPositiveCumulativeSum(arr []int) []int {
  result := []int{}
  tempSum := 0
  for _, num := range arr {
    tempSum += num
    if tempSum > 0 {
      result = append(result, tempSum)
    }
  }
  return result
}
