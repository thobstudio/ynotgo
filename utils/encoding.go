package utils

type Encoder struct {
	current_position int
	current_buffer   []byte
	buffers          [][]byte
}

func NewEncoder() *Encoder {
	return &Encoder{
		current_position: 0,
		current_buffer:   make([]byte, 100),
		buffers:          make([][]byte, 0),
	}
}

func (encoder *Encoder) Length() int {
	length := encoder.current_position
	for i := 0; i < len(encoder.buffers); i++ {
		length = length + len(encoder.buffers[i])
	}
	return length
}

func (encoder *Encoder) HasContent() bool {
	return encoder.current_position > 0 || len(encoder.buffers) > 0
}

func (encoder *Encoder) Write(num byte) {
	bufferLength := len(encoder.current_buffer)
	if encoder.current_position == bufferLength {
		encoder.buffers = append(encoder.buffers, encoder.current_buffer)
		encoder.current_buffer = make([]byte, bufferLength*2)
		encoder.current_position = 0
	}

	encoder.current_buffer[encoder.current_position] = num
	encoder.current_position++
}

func (encoder *Encoder) WriteVarUint(num uint) {
	for num > BITS7 {
		encoder.Write(byte(BIT8 | (BITS7 & num)))
		num = num / 128
	}
	encoder.Write(byte(BITS7 & num))
}

func (encoder *Encoder) ToUint8Array() []byte {
	uint8array := make([]byte, encoder.Length())
	current_position := 0
	for _, buffer := range encoder.buffers {
		for j, d := range buffer {
			uint8array[current_position+j] = d
		}
		current_position += len(buffer)
	}
	for i := 0; i < encoder.current_position; i++ {
		uint8array[current_position+i] = encoder.current_buffer[i]
	}
	return uint8array
}
