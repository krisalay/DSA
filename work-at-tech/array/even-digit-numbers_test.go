package workAtTech_Array

import (
  "testing"
  "reflect"
)

func TestGetEvenDigitNumbers(t *testing.T) {
  tests := []struct{
    name string
    input []int
    expected []int
  }{
    { "testcase 1", []int{3, 11, 4, 200}, []int{11} },
    { "testcase 2", []int{11, 42, 564, 5775, 34, 123, 454, 1, 5, 45, 3556, 23442}, []int{11, 42, 5775, 34, 45, 3556} },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := GetEvenDigitNumbers(tt.input)
      if !reflect.DeepEqual(result, tt.expected) {
        t.Errorf("GetEvenDigitNumbers(%v) = %v; expected %v", tt.input, result, tt.expected)
      }
    })
  }
}
