package rango

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	helper "rango/src/Helper"
)

func Decryption() ([]byte, error) {
	key := "TestingKey"
	// Now I gotta find a way to use the previous key to decrypt the file

	var original_data []byte

	ciphered_files, err := helper.Openfolder()
	helper.Error(err)

	for _, ciphered_file := range ciphered_files {
		// Read the encrypted data from file
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

		// Properly separate nonce and ciphertext
		nonce := encryptedData[:NonceSize]
		ciphertext := encryptedData[NonceSize:] // This line was missing

		// Decrypt the data using the correct ciphertext portion
		original_data, err := gcmInstance.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			return nil, fmt.Errorf("decryption failed: %v", err)
		}

		// Write the decrypted data back to file
		err = os.WriteFile(ciphered_file, original_data, 0644)
		helper.Error(err)

	}
	return original_data, nil
}
