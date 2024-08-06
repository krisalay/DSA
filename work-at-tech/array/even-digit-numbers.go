package workAtTech_Array

func HasEvenDigits(a int) bool {
	count := 0
	for a != 0 {
		a = a / 10
		count++
	}
	return count%2 == 0
}

func GetEvenDigitNumbers(arr []int) []int {
  result := []int{}

  for _, num := range arr {
    hasEven := HasEvenDigits(num)
    if hasEven == true {
      result = append(result, num)
    }
  }

  return result
}
