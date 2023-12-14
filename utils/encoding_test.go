package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	tests := []map[string][]uint{
		{"in": []uint{0}, "out": []uint{0}},
		{"in": []uint{1}, "out": []uint{1}},
		{"in": []uint{128}, "out": []uint{128, 1}},
		{"in": []uint{200}, "out": []uint{200, 1}},
		{"in": []uint{32}, "out": []uint{32}},
		{"in": []uint{500}, "out": []uint{244, 3}},
		{"in": []uint{256}, "out": []uint{128, 2}},
		{"in": []uint{700}, "out": []uint{188, 5}},
		{"in": []uint{1024}, "out": []uint{128, 8}},
		{"in": []uint{1025}, "out": []uint{129, 8}},
		{"in": []uint{4048}, "out": []uint{208, 31}},
		{"in": []uint{5050}, "out": []uint{186, 39}},
		{"in": []uint{1000000}, "out": []uint{192, 132, 61}},
		{"in": []uint{34951959}, "out": []uint{151, 166, 213, 16}},
		{"in": []uint{2147483646}, "out": []uint{254, 255, 255, 255, 7}},
		{"in": []uint{2147483647}, "out": []uint{255, 255, 255, 255, 7}},
		{"in": []uint{2147483648}, "out": []uint{128, 128, 128, 128, 8}},
		{"in": []uint{2147483700}, "out": []uint{180, 128, 128, 128, 8}},
		{"in": []uint{4294967294}, "out": []uint{254, 255, 255, 255, 15}},
		{"in": []uint{4294967295}, "out": []uint{255, 255, 255, 255, 15}},
	}
	for _, test := range tests {
		encoder := NewEncoder()
		encoder.WriteVarUint(test["in"][0])
		buffer := encoder.ToUint8Array()
		assert.Equal(t, len(buffer), len(test["out"]))
		assert.Greater(t, len(buffer), 0)
		assert.True(t, encoder.HasContent())
		for i := 0; i < len(buffer); i++ {
			assert.Equal(t, buffer[i], byte(test["out"][i]))
		}
	}
}

func TestVarUint(t *testing.T) {
	scenarios := []uint{42, 1<<9 | 3, 1<<17 | 1<<9 | 3, 1<<25 | 1<<17 | 1<<9 | 3, 2839012934}
	for _, scenario := range scenarios {
		encoder := NewEncoder()
		encoder.WriteVarUint(uint(scenario))
		buffer := encoder.ToUint8Array()
		decoder := NewDecoder(buffer)
		result, err := decoder.ReadVarUint()
		assert.Nil(t, err)
		assert.Equal(t, scenario, result)
	}
}
