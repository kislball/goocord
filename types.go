package goocord

// Base represents something that can be got from a primitive
type Base interface {
	FromPrimitive(struct{}) // FromPrimitive forces structure to base on primitive
	ToPrimitive() struct{} // ToPrimitive converts structure to primitive
}
