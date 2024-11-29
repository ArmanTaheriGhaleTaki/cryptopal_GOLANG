package set1

import (
	set1 "cryptopal_GOLANG/set1"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecryptTheMessage(t *testing.T) {
	const hexString string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	const expected string = "Cooking MC's like a pound of bacon"

	output, err := set1.DecryptTheMessage(hexString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.Equal(t, output, []byte(expected), "The two strings should be the same.")

}
