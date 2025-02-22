package logic2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal4b(t *testing.T) {
	result := Soal4b(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[0][0], 1)
	assert.Equal(t, result[1][7], 49)
}
