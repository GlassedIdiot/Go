package main

import (
	"fmt"
	"hats/Stegnography/PngLib"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		os.Exit(1)
	}

	// Get the file path from the command line arguments
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file '%s': %v", filePath, err)
	}
	defer file.Close()

	// Preprocess the file to get a bytes.Reader
	bReader, err := Steganography.PreProcessing(file)
	if err != nil {
		log.Fatalf("Preprocessing failed: %v", err)
	}

	header := &Steganography.Header{}

	// Validate the PNG header
	header.ValidatePNG(bReader)
}
