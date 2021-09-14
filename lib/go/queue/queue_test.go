package queue

import (
	"testing"
)

func TestBasicQueue(t *testing.T) {
	exp := NewQUEUE()
	ref := make([]DATATYPE,0)

	for k := 0; k < 2; k++ {
		for j := 0; j < 10; j++ {
			for i := 0; i < 16; i++ {
				ref = append([]DATATYPE{DATATYPE(i)},ref...)
				exp.Push(DATATYPE(i))
				if exp.Len() != len(ref) { t.Error("ERROR 1"); return }
				if exp.Tail() != ref[len(ref)-1] { t.Error("ERROR 2"); return }
				if exp.Len() != len(ref) { t.Error("ERROR 3"); return }
				if exp.IsEmpty() != (len(ref)==0) { t.Error("ERROR 4"); return }
			}
			for i := 15; i >= 0; i-- {
				v1 := ref[len(ref)-1]; ref = ref[:len(ref)-1]
				v2 := exp.Pop()
				if v1 != v2 { t.Error("ERROR 5"); return }
				if i > 0 {
					if exp.Head() != ref[0] { t.Error("ERROR 6"); return }
					if exp.Tail() != ref[len(ref)-1] { t.Error("ERROR 7"); return }
				}
				if exp.Len() != len(ref) { t.Error("ERROR 8"); return }
				if exp.IsEmpty() != (len(ref)==0) { t.Error("ERROR 9"); return }
			}
		}
		exp.Clear()
		ref = make([]DATATYPE,0)
	}
}
