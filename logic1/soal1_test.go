package logic1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	result := Soal1(9)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[3], 7)
	assert.NotEqual(t, result[3], 9)

}
