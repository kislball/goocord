package utils

// Flags represents a bitfield
type Flags struct {
	Flags int // Bitfield
}

// Add adds a new bit
func (f *Flags) Add(flags ...int) {
	for _, flag := range flags {
		f.Flags = f.Flags | flag
	}
}

// Has checks if a bit equals one
func (f *Flags) Has(flag int) bool {
	return f.Flags&flag == flag
}
