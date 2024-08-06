package workAtTech_Array

import "fmt"

func GetIdenticalTwinsCount(arr []int) int {
  result := 0
  countMap := map[int]int{}

  for _, num := range arr {
    if val, ok := countMap[num]; !ok {
      countMap[num] = 1
    } else {
      countMap[num] = val + 1
    }
  }

  for k, v := range countMap {
    fmt.Println(k, v)
    result += (v * (v-1)/2)
  }
  return result
}
