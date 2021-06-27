package goocord

// Base represents something that can be got from a primitive
type Base interface {
	FromPrimitive(struct{})
	ToPrimitive() struct{}
}
