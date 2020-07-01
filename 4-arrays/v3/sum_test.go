package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

