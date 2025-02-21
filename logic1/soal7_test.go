package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal7(t *testing.T) {
	result := Soal71(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[2], 5)
	assert.Equal(t, result[3], 7)
	assert.Equal(t, result[4], 9)
	assert.Equal(t, result[5], 9)
	assert.NotEqual(t, result[6], 5)
	assert.NotEqual(t, result[7], 3)

	result = Soal72(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[5], 11)
	assert.Equal(t, result[6], 9)
	assert.Equal(t, result[7], 7)
}
