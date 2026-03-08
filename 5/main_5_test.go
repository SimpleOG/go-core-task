package main

import (
	"reflect"
	"sort"
	"testing"
)

func sortedInts(s []int) []int {
	cp := make([]int, len(s))
	copy(cp, s)
	sort.Ints(cp)
	return cp
}

func TestFindIntersections_ExampleFromTZ(t *testing.T) {
	// a=[65,3,58,678,64], b=[64,2,3,43] → true, [64,3]
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	gotBool, gotSlice := FindIntersections(a, b)

	if !gotBool {
		t.Errorf("expected true, got false")
	}
	want := []int{3, 64}
	if !reflect.DeepEqual(sortedInts(gotSlice), want) {
		t.Errorf("FindIntersections() = %v, want %v (any order)", gotSlice, want)
	}
}

func TestFindIntersections_DuplicatesInArr1_CountedOnce(t *testing.T) {
	// a=[1,1,1] - значение 1 встречается трижды в arr1
	// результат должен содержать 1 ровно один раз
	a := []int{1, 1, 1}
	b := []int{1, 2, 3}

	gotBool, gotSlice := FindIntersections(a, b)

	if !gotBool {
		t.Errorf("expected true, got false")
	}
	if !reflect.DeepEqual(gotSlice, []int{1}) {
		t.Errorf("FindIntersections() = %v, want [1]", gotSlice)
	}
}

func TestFindIntersections_DuplicatesInArr2_CountedOnce(t *testing.T) {
	// значение 1 встречается трижды в arr2 - в результате должно быть один раз
	a := []int{1, 2, 3}
	b := []int{1, 1, 1}

	gotBool, gotSlice := FindIntersections(a, b)

	if !gotBool {
		t.Errorf("expected true, got false")
	}
	if !reflect.DeepEqual(gotSlice, []int{1}) {
		t.Errorf("FindIntersections() = %v, want [1]", gotSlice)
	}
}

func TestFindIntersections_DuplicatesInBoth_CountedOnce(t *testing.T) {
	a := []int{1, 1, 2, 2}
	b := []int{1, 1, 2, 2}

	gotBool, gotSlice := FindIntersections(a, b)

	if !gotBool {
		t.Errorf("expected true, got false")
	}
	want := []int{1, 2}
	if !reflect.DeepEqual(sortedInts(gotSlice), want) {
		t.Errorf("FindIntersections() = %v, want %v", gotSlice, want)
	}
}

func TestFindIntersections_NoIntersections_ReturnsFalseAndEmptySlice(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{61, 2, 4, 43}

	gotBool, gotSlice := FindIntersections(a, b)

	if gotBool {
		t.Errorf("expected false, got true")
	}
	if len(gotSlice) != 0 {
		t.Errorf("expected empty slice, got %v", gotSlice)
	}
}

func TestFindIntersections_EmptyFirstSlice_ReturnsFalseAndEmpty(t *testing.T) {
	gotBool, gotSlice := FindIntersections([]int{}, []int{1, 2, 3})

	if gotBool {
		t.Errorf("expected false, got true")
	}
	if len(gotSlice) != 0 {
		t.Errorf("expected empty slice, got %v", gotSlice)
	}
}

func TestFindIntersections_EmptySecondSlice_ReturnsFalseAndEmpty(t *testing.T) {
	gotBool, gotSlice := FindIntersections([]int{1, 2, 3}, []int{})

	if gotBool {
		t.Errorf("expected false, got true")
	}
	if len(gotSlice) != 0 {
		t.Errorf("expected empty slice, got %v", gotSlice)
	}
}

func TestFindIntersections_BothEmpty_ReturnsFalseAndEmpty(t *testing.T) {
	gotBool, gotSlice := FindIntersections([]int{}, []int{})

	if gotBool {
		t.Errorf("expected false, got true")
	}
	if len(gotSlice) != 0 {
		t.Errorf("expected empty slice, got %v", gotSlice)
	}
}

func TestFindIntersections_AllElementsMatch(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	gotBool, gotSlice := FindIntersections(a, b)

	if !gotBool {
		t.Errorf("expected true, got false")
	}
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(sortedInts(gotSlice), want) {
		t.Errorf("FindIntersections() = %v, want %v", gotSlice, want)
	}
}

func TestFindIntersections_NegativeNumbers(t *testing.T) {
	a := []int{-1, -2, -3, 4}
	b := []int{-2, 5, -3}

	gotBool, gotSlice := FindIntersections(a, b)

	if !gotBool {
		t.Errorf("expected true, got false")
	}
	want := []int{-3, -2}
	if !reflect.DeepEqual(sortedInts(gotSlice), want) {
		t.Errorf("FindIntersections() = %v, want %v", gotSlice, want)
	}
}

func TestFindIntersections_BoolConsistentWithSlice(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
	}{
		{"has intersections", []int{1, 2, 3}, []int{2, 4}},
		{"no intersections", []int{1, 2, 3}, []int{4, 5, 6}},
		{"both empty", []int{}, []int{}},
		{"duplicates in arr1", []int{1, 1, 1}, []int{1}},
		{"duplicates in arr2", []int{1}, []int{1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotSlice := FindIntersections(tt.a, tt.b)
			if gotBool != (len(gotSlice) > 0) {
				t.Errorf("bool=%v не соответствует len(slice)=%d", gotBool, len(gotSlice))
			}
		})
	}
}
