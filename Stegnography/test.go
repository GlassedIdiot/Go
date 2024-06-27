package main

import (
	"encoding/hex"
)

func main() {
	src := []byte("89504e47")
	dst := make([]byte, hex.DecodedLen(len(src)))
}
