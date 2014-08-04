package data

import "testing"

func TestCompare(t *testing.T) {
	iw1 := &IntWeight{1}
	iw2 := &IntWeight{2}
	iw3 := &IntWeight{3}

	if iw2.Compare(iw2) != 0 {
		t.Error("Same IntWeight's Compare result should be 0")
	}

	if iw2.Compare(iw1) != 1 {
		t.Error("IntWeight with value 2 should return 1(greater) in comparsion with value 1")
	}

	if iw2.Compare(iw3) != -1 {
		t.Error("IntWeight with value 2 should return 1(greater) in comparsion with value 3")
	}
}
