package set1

import (
	hex "encoding/hex"
	"regexp"
	"strings"
)

var letterFrequencies = map[string]float64{
	"A": 8.167,
	"B": 1.492,
	"C": 2.782,
	"D": 4.253,
	"E": 12.702,
	"F": 2.228,
	"G": 2.015,
	"H": 6.094,
	"I": 6.966,
	"J": 0.153,
	"K": 0.772,
	"L": 4.025,
	"M": 2.406,
	"N": 6.749,
	"O": 7.507,
	"P": 1.929,
	"Q": 0.095,
	"R": 5.987,
	"S": 6.327,
	"T": 9.056,
	"U": 2.758,
	"V": 0.978,
	"W": 2.360,
	"X": 0.150,
	"Y": 1.974,
	"Z": 0.074,
}
var text = regexp.MustCompile("^[a-zA-Z ]$")

func isAlphabetic(str string) bool {
	return text.MatchString(str)
}

func calculateScore(buffer []byte) float64 {
	score := 0.0
	for _, b := range buffer {
		str := string(b)
		if isAlphabetic(str) {
			score += letterFrequencies[strings.ToUpper(str)]
		} else {
			score -= 10.0
		}
	}

	return score
}

func singleByteXOR(buffer []byte, key byte) []byte {
	result := make([]byte, len(buffer))
	for i := range len(buffer) {
		result[i] = buffer[i] ^ key
	}

	return result
}
func DecryptTheMessage(hexString string) ([]byte, error) {
	cipher, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	bestKey := byte(0)
	bestScore := 0.0
	for k := range 256 {
		key := byte(k)
		decrypted := singleByteXOR(cipher, key)

		score := calculateScore(decrypted)

		if score > bestScore {
			bestScore = score
			bestKey = key
		}
	}

	decryptedMessage := singleByteXOR(cipher, bestKey)
	return decryptedMessage, nil
}
