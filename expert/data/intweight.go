package data

//IntWeight is the common integer weight interface implementation
type IntWeight struct {
	Value int
}

//Compare is the Weight interface function implementation
func (iw *IntWeight) Compare(w Weight) int {
	iw1 := w.(*IntWeight)

	if iw.Value == iw1.Value {
		return 0
	} else if iw.Value < iw1.Value {
		return -1
	} else {
		return 1
	}
}
