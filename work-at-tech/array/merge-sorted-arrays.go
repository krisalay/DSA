package workAtTech_Array

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
  size1 := len(arr1)
  size2 := len(arr2)

  if size1 == 0 && size2 == 0 {
    return []int{}
  }

  if size1 == 0 {
    return arr2
  }
  if size2 == 0 {
    return arr1
  }

  p1, p2 := 0, 0
  result := []int{}
  for p1 < size1 && p2 < size2 {
    if arr1[p1] == arr2[p2] {
      result = append(result, arr1[p1])
      result = append(result, arr2[p2])
      p1++
      p2++
    } else if arr1[p1] > arr2[p2] {
      result = append(result, arr2[p2])
      p2++
    } else {
      result = append(result, arr1[p1])
      p1++
    }
  }

  for p1 < size1 {
    result = append(result, arr1[p1])
    p1++
  }

  for p2 < size2 {
    result = append(result, arr2[p2])
    p2++
  }

  return result
}
