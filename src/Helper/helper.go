package helper

import (
	"crypto/rand"
	"log"
	"os"
	"path/filepath"
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
func Openfolder() ([]string, error) {
	files := []string{}
	err := filepath.WalkDir("D:\\Shit\\Go\\Test_folder", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only process if it's not a directory
		if !d.IsDir() {
			// Read the file contents

			println(path)
			files = append(files, path)

		}
		return nil
	})
	Error(err)
	return files, err
}
