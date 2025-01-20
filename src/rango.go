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

func Openfolder() {
	err := filepath.WalkDir("D:\\Shit\\Go\\Test_folder", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Only process if it's not a directory
		if !d.IsDir() {
			// Read the file contents

			println(path)
			//Encryption(path)
			Decryption(path)
		}
		return nil
	})
	helper.Error(err)
}

func Encryption(test_file string) ([]byte, error) {

	key := "TestingKey"

	file_content, err := os.ReadFile(test_file)

	if err != nil {
		helper.Error(err)
	}
	fmt.Printf("%s\n and the size of this file is %d", file_content, len(file_content))

	aesBlock, err := aes.NewCipher([]byte(mdHashing(string(key))))
	helper.Error(err)

	gcmInstance, err := cipher.NewGCM(aesBlock)

	helper.Error(err)

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	ciphered_text := gcmInstance.Seal(nonce, nonce, []byte(file_content), nil)

	fmt.Printf("\n%s", ciphered_text)

	err = os.WriteFile(test_file, ciphered_text, 0644)

	helper.Error(err)
	return nil, nil
}

func Decryption(ciphered_file string) ([]byte, error) {
	key := "TestingKey"
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

	// Print the decrypted data
	fmt.Printf("Decrypted content: %s\n", string(original_data))

	return original_data, nil
}
