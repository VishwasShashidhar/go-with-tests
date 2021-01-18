package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum slices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("sum all - two", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{5, 10, 15}

		got := SumAll(numbers1, numbers2)
		want := []int{15, 30}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("sum tail of a slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{5, 10, 15})
		want := []int{14, 25}

		checkSums(t, got, want)
	})

	t.Run("sum tail of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{})
		want := []int{14, 0}

		checkSums(t, got, want)
	})
}
