package logic2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal5b(t *testing.T) {
	result := Soal5b(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][4], 9)
	assert.Equal(t, result[1][2], 31)
	assert.NotEqual(t, result[0][1], 19)
}
