package data

import "testing"

func TestCompare(t *testing.T) {
	sw1 := &SingleWeight{1}
	sw2 := &SingleWeight{2}
	sw3 := &SingleWeight{3}

	if sw2.Less(sw2)  {
		t.Error("Same SingleWeight's Less result should be true")
	}

	if sw2.Less(sw1) {
		t.Error("SingleWeight with value 2 should return false(greater) in comparsion with value 1")
	}

	if !sw2.Less(sw3) {
		t.Error("SingleWeight with value 2 should return true(lesser) in comparsion with value 3")
	}
}

func TestAdd(t *testing.T) {
	sw1 := &SingleWeight{1}
	sw2 := &SingleWeight{2}
	
	sw1.Add(sw2)

	if sw1.Value.(int64) != 3 {
		t.Error("Add should do 1 + 2 == 3, not ", sw1.Value)
	}
}

func BenchmarkCommonCalculation(b *testing.B) {
	sw1 := &SingleWeight{1}
	sw2 := &SingleWeight{2}
	for i := 0; i < b.N; i++ {
		sw1.Less(sw2)
	}
}
