package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Error(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

func Openfolder() ([]string, error) {
	// Load variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	files := []string{}
	folderPath := os.Getenv("TEST_FOLDER_PATH")
	fmt.Printf("Folder Path: %s\n", folderPath)
	err = filepath.WalkDir(folderPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			log.Fatal("WalkDir error:", err)
		}

		// Only process if it's not a directory
		if !d.IsDir() {
			fmt.Println(path)
			files = append(files, path)
		}
		return nil
	})
	Error(err)
	return files, err
}
