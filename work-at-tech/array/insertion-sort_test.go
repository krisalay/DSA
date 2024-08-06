package workAtTech_Array

import (
  "testing"
  "reflect"
)

func TestInsertionSort(t *testing.T) {
  tests := []struct{
    name string
    input []int
    expected []int
  }{
    { "testcase 1", []int{1, 4, 5, 3, 2, 0}, []int{0, 1, 2, 3, 4, 5} },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      insertionSort(&(tt.input))
      t.Logf("After sorting: %v", tt.input)
      if !reflect.DeepEqual(tt.input, tt.expected) {
        t.Errorf("insertionSort(%v) = %v; expected %v", tt.input, tt.input, tt.expected)
      }
    })
  }
}
