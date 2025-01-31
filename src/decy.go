package rango

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"path/filepath"
	helper "rango/src/Helper"
)

func Decryption() ([]byte, error) {
	key := "TestingKey"
	ciphered_files, err := helper.Openfolder()
	helper.Error(err)

	for _, ciphered_file := range ciphered_files {
		// Skip if the file is not encrypted
		if filepath.Ext(ciphered_file) != ".enc" {
			fmt.Print("Skipping: ", ciphered_file, "\n")
			continue
		}

		encryptedData, err := os.ReadFile(ciphered_file)
		if err != nil {
			return nil, err
		}

		aesBlock, err := aes.NewCipher([]byte(mdHashing(string(key))))
		helper.Error(err)

		gcmInstance, err := cipher.NewGCM(aesBlock)
		helper.Error(err)

		NonceSize := gcmInstance.NonceSize()

		if len(encryptedData) < NonceSize {
			return nil, fmt.Errorf("ciphertext too short")
		}

		nonce := encryptedData[:NonceSize]
		ciphertext := encryptedData[NonceSize:]

		original_data, err := gcmInstance.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			return nil, fmt.Errorf("decryption failed: %v", err)
		}

		// Remove the .enc extension to restore the original filename
		originalFilePath := ciphered_file[:len(ciphered_file)-4]
		err = os.WriteFile(originalFilePath, original_data, 0644)
		helper.Error(err)

		// Optionally, remove the encrypted file
		err = os.Remove(ciphered_file)
		helper.Error(err)
	}
	return nil, nil
}
