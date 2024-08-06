package workAtTech_Array

import (
  "testing"
  "reflect"
)

func TestMergeSortedArrays(t *testing.T) {
  tests := []struct{
    name string
    arr1 []int
    arr2 []int
    expected []int
  }{
    {"testcase 1", []int{1, 2, 3}, []int{5, 8, 9}, []int{1, 2, 3, 5, 8, 9} },
  }

  for _, tt := range tests {
    t.Run(tt.name, func (t *testing.T) {
      result := mergeSortedArrays(tt.arr1, tt.arr2)
      if !reflect.DeepEqual(result, tt.expected) {
        t.Errorf("mergeSortedArrays(%v, %v) = %v; Expected = %v", tt.arr1, tt.arr2, result, tt.expected)
      }
    })
  }
}
