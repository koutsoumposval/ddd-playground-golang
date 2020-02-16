package value

// ID represents an identifier
type ID struct {
	id int64
}

// ID returns the id
func (identifier *ID) ID() int64 {
	return identifier.id
}

// SetID sets the id
func (identifier *ID) SetID(id int64) {
	identifier.id = id
}
