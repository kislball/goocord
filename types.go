package goocord

// Base represents something that can be got from a primitive
type Base interface {
	// Force structure to base on primitive
	FromPrimitive(struct{})
	// Convert structure to primitive
	ToPrimitive() struct{}
}
