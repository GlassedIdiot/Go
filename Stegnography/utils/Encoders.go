package utils

import "fmt"

func EncodeDecode(payload []byte, xorkey string) []byte {
	de_byte := make([]byte, len(xorkey))

	for i := 0; i < len(xorkey); i++ {
		de_byte[i] = payload[i] ^ xorkey[i%len(xorkey)]
	}

	fmt.Printf(
		"The encoded payload as a bytes array is: %d.\nThe Payload encoded as a string is: %s\n",
		de_byte,
		de_byte,
	)
	return nil
}
