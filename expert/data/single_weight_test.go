package data

import "testing"

func TestCompare(t *testing.T) {
	iw1 := &SingleWeight{1}
	iw2 := &SingleWeight{2}
	iw3 := &SingleWeight{3}

	if iw2.Compare(iw2) != 0 {
		t.Error("Same SingleWeight's Compare result should be 0")
	}

	if iw2.Compare(iw1) != 1 {
		t.Error("SingleWeight with value 2 should return 1(greater) in comparsion with value 1")
	}

	if iw2.Compare(iw3) != -1 {
		t.Error("SingleWeight with value 2 should return 1(greater) in comparsion with value 3")
	}
}

func BenchmarkCommonCalculation(b *testing.B) {
	iw1 := &SingleWeight{1}
	iw2 := &SingleWeight{2}
	for i := 0; i < b.N; i++ {
		iw1.Compare(iw2)
	}
}
