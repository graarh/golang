package data

// Weight is an abstract interface for any weight realisation
type Weight interface {
	Compare(w Weight) int
}
