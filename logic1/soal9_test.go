package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal9(t *testing.T) {
	result := Soal91(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[4], 15)
	assert.NotEqual(t, result[5], 12)

	result = Soal92(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[5], 18)
	assert.NotEqual(t, result[6], 18)
}
