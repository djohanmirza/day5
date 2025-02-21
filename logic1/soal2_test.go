package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal2(t *testing.T) {
	result := Soal2(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[3], 8)
	assert.NotEqual(t, result[3], 12)

}
