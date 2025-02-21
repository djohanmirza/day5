package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal8(t *testing.T) {
	result := Soal81(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[4], 10)
	assert.Equal(t, result[7], 6)
	assert.NotEqual(t, result[5], 8)

	result = Soal82(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[5], 12)
	assert.Equal(t, result[6], 10)
	assert.NotEqual(t, result[8], 8)
}
