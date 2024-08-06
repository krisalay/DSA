package workAtTech_Array

import (
  "testing"
  "reflect"
)

func TestRotateMatrix(t *testing.T) {
  tests := []struct{
    name string
    input [][]int
    expected [][]int
  }{
    { "testcase1", [][]int{ {1, 2}, {3, 4}, {5, 6} }, [][]int{ {5, 3, 1}, { 6, 4, 2 } } },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T){
      result := rotateMatrix(tt.input)
      t.Logf("After rotation: %v", result)
      if !reflect.DeepEqual(result, tt.expected) {
        t.Errorf("rotateMatrix(%v) = %v; Expected = %v", tt.input, result, tt.expected)
      }
    })
  }

}
