package helper

import (
	"crypto/rand"
	"log"
)

func Error(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
func GenerateCryptoRandom(chars string, length int32) string {
	bytes := make([]byte, length)
	rand.Read(bytes)

	for index, element := range bytes {
		randomize := element % byte(len(chars))
		bytes[index] = chars[randomize]
	}
	return string(bytes)
}
