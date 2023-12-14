package utils

import "errors"

type Decoder struct {
	bytes    []byte
	position int
}

func NewDecoder(bytes []byte) *Decoder {
	return &Decoder{
		bytes:    bytes,
		position: 0,
	}
}

func (decoder *Decoder) HasContent() bool {
	return decoder.position != len(decoder.bytes)
}

func (decoder *Decoder) ReadVarUint() (uint, error) {
	num := uint(0)
	mult := uint(1)
	length := len(decoder.bytes)
	for decoder.position < length {
		r := decoder.bytes[decoder.position]
		decoder.position++
		num = num + (uint(r)&BITS7)*mult
		mult *= 128
		if r < BIT8 {
			return num, nil
		}
	}

	return 0, errors.New("unexpected end of array")
}
