package workAtTech_Array

func insertionSort(arr *[]int) {
  size := len(*arr)
  if size <= 1 {
    return
  }

  write, read := 0, 1

  for read < size {
    if (*arr)[read] < (*arr)[write] {
      w, r := write, read
      for r >= 0 && w >= 0 && (*arr)[r] < (*arr)[w] {
        (*arr)[r], (*arr)[w] = (*arr)[w], (*arr)[r]
        r--
        w--
      }
    }
    write++
    read++
  }
}
