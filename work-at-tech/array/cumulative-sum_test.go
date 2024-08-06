package workAtTech_Array

import (
  "testing"
  "reflect"
)

func TestCumulativeSum(t *testing.T) {
  tests := []struct {
    name string
    input []int
    expected []int
  }{
    {"cumulative sum", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := GetCumulativeSum(tt.input)
      if !reflect.DeepEqual(result, tt.expected) {
        t.Errorf("GetCumulativeSum(%v) = %v; expected %v", tt.input, result, tt.expected)
      }
    })
  }
}
