package goocord

// Base represents something that can be got from a primitive
type Base interface {
	FromPrimitive(interface{}) // FromPrimitive forces structure to base on primitive
	ToPrimitive() interface{}  // ToPrimitive converts structure to primitive
}
