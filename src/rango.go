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

			//Encryption(path)
			Decryption(path)
		}
		return nil
	})
	helper.Error(err)
}

func Encryption(test_file string) ([]byte, error) {

	key := "TestingKey"

	_, err := os.ReadFile(test_file)

	if err != nil {
		helper.Error(err)
	}
	fmt.Printf("%s\n", test_file)

	fmt.Printf("%s", test_file)

	aesBlock, err := aes.NewCipher([]byte(mdHashing(string(key))))
	helper.Error(err)

	gcmInstance, err := cipher.NewGCM(aesBlock)

	helper.Error(err)

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	ciphered_text := gcmInstance.Seal(nonce, nonce, []byte(test_file), nil)

	fmt.Printf("%s", ciphered_text)

	err = os.WriteFile(test_file, ciphered_text, 0644)

	helper.Error(err)
	return ciphered_text, nil
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

	// Make sure the encrypted data is large enough to contain the nonce
	if len(encryptedData) < NonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Extract nonce and ciphertext from the encrypted data
	nonce, ciphered_Data := encryptedData[:NonceSize], encryptedData[NonceSize:]

	// Decrypt the data
	original_file, err := gcmInstance.Open(nil, nonce, ciphered_Data, nil)

	err = os.WriteFile(ciphered_file, original_file, 0644)
	helper.Error(err)

	return original_file, nil
}
