package main

import (
	"fmt"
	"hats/Stegnography/PngLib"
	"hats/Stegnography/utils"
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

	meta_Cnk := &Steganography.MetaChunk{}
	// Validate the PNG header
	meta_Cnk.ValidatePNG(bReader)
	// meta_Cnk.ParsePNG(bReader, nil)

	payload := []byte("fmhy.net,goland.net,blackchair")
	xorkey := "findme"

	fmt.Printf("%s\n", payload)

	something := utils.XorEncode(payload, xorkey)

	fmt.Printf("%s\n", something)

	something_test := utils.XorDecode(something, xorkey)

	fmt.Printf("%s\n", something_test)
}
