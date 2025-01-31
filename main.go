package main

import (
	"fmt"
	"path/filepath"
	rango "rango/src"
	helper "rango/src/Helper"
)

func main() {
	files, err := helper.Openfolder() // Adjust if needed
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	allEncrypted := true
	for _, f := range files {
		if filepath.Ext(f) != ".enc" {
			allEncrypted = false
			break
		}
	}

	if allEncrypted {
		fmt.Print("Files are already encrypted. Would you like to decrypt them? (y/n): ")
		var decryptOpt string
		fmt.Scanln(&decryptOpt)
		if decryptOpt == "y" || decryptOpt == "yes" {
			_, decErr := rango.Decryption()
			if decErr != nil {
				fmt.Println("Error decrypting files:", decErr)
			} else {
				fmt.Println("Files decrypted successfully.")
			}
		}
	} else {
		fmt.Print("Would you like to encrypt the files? (y/n): ")
		var encryptOpt string
		fmt.Scanln(&encryptOpt)
		if encryptOpt == "y" || encryptOpt == "yes" {
			_, encErr := rango.Encryption()
			if encErr != nil {
				fmt.Println("Error encrypting files:", encErr)
			} else {
				fmt.Println("Files encrypted successfully.")
			}
		}
	}
}
