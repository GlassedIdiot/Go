package utils

func EncodeDecode(payload []byte, key string) []byte {
	de_bytes := make([]byte, len(payload))
	for i := 0; i < len(payload); i++ {
		de_bytes[i] = payload[i] ^ key[i%len(key)]
	}
	return de_bytes
}

// XorEncode returns encoded byte array
func XorEncode(Decode []byte, key string) []byte {
	return EncodeDecode(Decode, key)
}

// XorDecode returns decoded byte array
func XorDecode(Encode []byte, key string) []byte {
	return EncodeDecode(Encode, key)
}
