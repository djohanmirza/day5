package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal4(t *testing.T) {
	result := Soal4(9)

	assert.Equal(t, len(result), 9)
	assert.NotEqual(t, len(result), 8)
	assert.Equal(t, result[5], 9)
	assert.NotEqual(t, result[4], 8)
}
