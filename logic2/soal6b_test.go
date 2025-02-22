package logic2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal6b(t *testing.T) {
	result := Soal6b(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][8], 25)
	assert.Equal(t, result[1][8], 28)
	assert.Equal(t, result[1][0], 44)
}
