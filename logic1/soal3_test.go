package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal3(t *testing.T) {
	result := Soal3(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[5], 18)
	assert.NotEqual(t, len(result), 8)
}
