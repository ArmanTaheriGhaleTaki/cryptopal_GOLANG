package set1

import (
	b64 "encoding/base64"
	hex "encoding/hex"
	"log"
)

func Convert_hex_to_base64(input string) string {
	temp, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	encoded := b64.StdEncoding.EncodeToString(temp)
	return encoded

}
