package utils

type ID struct {
	client uint32
	clock  uint32
}

func NewID(client uint32, clock uint32) *ID {
	return &ID{
		client: client,
		clock:  clock,
	}
}

func (id *ID) CompareID(other *ID) bool {
	return id == other ||
		(id != nil && other != nil && id.client == other.client && id.clock == other.clock)
}
