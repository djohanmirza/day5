package logic2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal2b(t *testing.T) {
	result := Soal2b(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[6][8], 18)
	assert.Equal(t, result[6][5], 12)
	assert.NotEqual(t, result[2][5], 14)
}
