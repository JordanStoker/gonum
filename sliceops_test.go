package sliceops

import (
	"testing"
)

const (
	SmallBenchmark  = 10
	MediumBenchmark = 1000
	LargeBenchmark  = 100000
)

func TestMin(t *testing.T) {
	s := []float64{}
	val, ind := Min(s)
	if val != 0 {
		t.Errorf("Val not returned as default when slice length is zero")
	}
	if ind != -1 {
		t.Errorf("Ind not returned as -1 for empty slice")
	}
	s = []float64{3, 4, 1, 7, 5}
	val, ind = Min(s)
	if val != 1 {
		t.Errorf("Wrong value returned")
	}
	if ind != 2 {
		t.Errorf("Wrong index returned")
	}
}

func TestMax(t *testing.T) {
	s := []float64{}
	val, ind := Max(s)
	if val != 0 {
		t.Errorf("Val not returned as default when slice length is zero")
	}
	if ind != -1 {
		t.Errorf("Ind not returned as -1 for empty slice")
	}
	s = []float64{3, 4, 1, 7, 5}
	val, ind = Max(s)
	if val != 7 {
		t.Errorf("Wrong value returned")
	}
	if ind != 3 {
		t.Errorf("Wrong index returned")
	}
}

func TestSum(t *testing.T) {
	s := []float64{}
	val := Sum(s)
	if val != 0 {
		t.Errorf("Val not returned as default when slice length is zero")
	}
	s = []float64{3, 4, 1, 7, 5}
	val = Sum(s)
	if val != 20 {
		t.Errorf("Wrong sum returned")
	}
}

func TestProd(t *testing.T) {
	s := []float64{}
	val := Prod(s)
	if val != 1 {
		t.Errorf("Val not returned as default when slice length is zero")
	}
	s = []float64{3, 4, 1, 7, 5}
	val = Prod(s)
	if val != 420 {
		t.Errorf("Wrong prod returned. Expected %v returned %v", 420, val)
	}
}

func TestHasEqLen(t *testing.T) {
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{1, 2, 3, 4}
	s3 := []float64{1, 2, 3}
	if !HasEqLen(s1, s2) {
		t.Errorf("Equal lengths returned as unequal")
	}
	if HasEqLen(s1, s3) {
		t.Errorf("Unequal lengths returned as equal")
	}
	if !HasEqLen(s1) {
		t.Errorf("Single slice returned as unequal")
	}
	if !HasEqLen() {
		t.Errorf("No slices returned as unequal")
	}
}
