package logic2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1b(t *testing.T) {
	result := Soal1b(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[8][8], 17)
	assert.Equal(t, result[3][4], 9)
	assert.NotEqual(t, result[0][4], 3)
}
