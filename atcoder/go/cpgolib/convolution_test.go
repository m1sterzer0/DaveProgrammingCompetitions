package cpgolib

import (
	"fmt"
	"testing"
)

func TestNTTSimple(t *testing.T) {
	eng := NewConvolver(998244353, 31, 23)
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	eng.NTT(a1, false)
	eng.NTT(a1, true)
	for i, a := range a1 {
		if a != i+1 {
			t.Error("NTT problem in TestNTTSimple")
			break
		}
	}
}

func TestConvolutionSimple(t *testing.T) {
	a := []int{1, 3, 3, 1}
	b := []int{1, 4, 6, 4, 1}
	eng := NewConvolver(998244353, 31, 23)
	c := eng.Convolve(a, b)
	cref := []int{1, 7, 21, 35, 35, 21, 7, 1}
	if len(c) != len(cref) {
		t.Error("Mismatched lengths")
	}
	for i := 0; i < len(c); i++ {
		if c[i] != cref[i] {
			s := fmt.Sprintf("Mismatched values c:%v cref:%v", c, cref)
			t.Error(s)
			break
		}
	}
}
