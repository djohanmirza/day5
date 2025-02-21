package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal6(t *testing.T) {
	result := Soal6(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[8], 6)
	assert.NotEqual(t, result[4], 27)
}
