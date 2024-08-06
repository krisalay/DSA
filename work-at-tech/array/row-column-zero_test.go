package workAtTech_Array

import (
  "testing"
  "reflect"
)

func TestSetRowColumnZeroes(t *testing.T) {
  tests := []struct {
    name string
    input [][]int
    expected [][]int
  }{
    { "testcase 1", [][]int{ {1, 1, 1}, {1, 0, 1}, {1, 1, 1} }, [][]int{ {1, 0, 1}, {0, 0, 0}, {1, 0, 1} } },
  }

  for _, tt := range tests {
    t.Run(tt.name, func (t *testing.T) {
      setRowColumnZeroes(&(tt.input))
      if !reflect.DeepEqual(tt.input, tt.expected) {
        t.Errorf("setRowColumnZeroes(%v); Expected = %v", tt.input, tt.expected)
      }
    })
  }
}
