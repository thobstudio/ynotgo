package utils

type ID struct {
	client uint
	clock  uint
}

func newID(client uint, clock uint) *ID {
	return &ID{
		client: client,
		clock:  clock,
	}
}
func NewID(client uint, clock uint) *ID {
	return newID(client, clock)
}

func (id *ID) CompareID(other *ID) bool {
	return id == other ||
		(id != nil && other != nil && id.client == other.client && id.clock == other.clock)
}

func (id *ID) WriteID(encoder *Encoder) {
	encoder.WriteVarUint(id.client)
	encoder.WriteVarUint(id.clock)
}

func ReadID(decoder *Decoder) (*ID, error) {
	client, err := decoder.ReadVarUint()
	if err != nil {
		return nil, err
	}
	clock, err := decoder.ReadVarUint()
	if err != nil {
		return nil, err
	}
	return newID(client, clock), nil
}
