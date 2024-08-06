package workAtTech_Array

import (
  "testing"
  "reflect"
)

func TestGetPositiveCumulativeSum(t *testing.T) {
  tests := []struct {
    name string
    input []int
    expected []int
  }{
    {"positive cumulative sum", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
    {"positive cumulative sum", []int{1, -2, 3, 4, 6}, []int{1, 2, 6, 12}},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := GetPositiveCumulativeSum(tt.input)
      if !reflect.DeepEqual(result, tt.expected) {
        t.Errorf("GetPositiveCumulativeSum(%v) = %v; expected %v", tt.input, result, tt.expected)
      }
    })
  }
}

