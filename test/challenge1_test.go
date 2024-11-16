package set1

import (
	"testing"

	set1 "cryptopal_GOLANG/set1"

	"github.com/stretchr/testify/assert"
)

func TestHex_To_Base64(t *testing.T) {
	const input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	output := set1.Convert_hex_to_base64(input)

	assert.Equal(t, output, expected, "The two strings should be the same.")

}
