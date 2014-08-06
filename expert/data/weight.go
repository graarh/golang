package data

// Weight is an abstract interface for any weight realisation
type Weight interface {
	Less(w Weight) bool
	Add(w Weight)
}
