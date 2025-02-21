package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal5(t *testing.T) {
	result := Soal5(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[4], 12)
	assert.NotEqual(t, result[4], 18)
}
