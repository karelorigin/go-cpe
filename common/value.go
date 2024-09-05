package common

// Value represents a
type Value struct {
	Quoted string
	Raw    string
}

// String always return the quoted value representation.
func (v Value) String() string {
	return v.Quoted
}
