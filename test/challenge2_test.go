package set1

import (
	"testing"

	set1 "cryptopal_GOLANG/set1"

	"github.com/stretchr/testify/assert"
)

func TestXor(t *testing.T) {
	const input1 string = "1c0111001f010100061a024b53535009181c"
	const input2 string = "686974207468652062756c6c277320657965"
	const expected string = "746865206b696420646f6e277420706c6179"
	output := set1.Xor(input1,input2)

	assert.Equal(t, output, expected, "The two strings should be the same.")

}
