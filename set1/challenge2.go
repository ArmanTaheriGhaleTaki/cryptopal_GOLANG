package set1

import (
	hex "encoding/hex"
)

func Xor(input1 string, input2 string) string {
	temp1, _ := hex.DecodeString(input1)
	temp2, _ := hex.DecodeString(input2)
	output := make([]byte, len(temp1))
	for i := range temp1 {
		output[i] = temp1[i] ^ temp2[i]
	}
	return hex.EncodeToString(output)

}
