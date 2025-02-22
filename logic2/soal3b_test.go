package logic2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal3b(t *testing.T) {
	result := Soal3b(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[3][7], 69)
	assert.Equal(t, result[1][1], 21)
}
