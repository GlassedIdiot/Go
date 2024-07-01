package Steganography

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func OpenPNGfile() (*os.File, error) {
	file, err := os.Open("/home/cindy/Pictures/wallpapers/TheMistyFields.png")
	checkErr(err)
	return file, err
}

func ValidatePNG(file *os.File) {
	signature := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	checkBytes := make([]byte, len(signature))
	n, err := file.Read(checkBytes)
	checkErr(err)

	if string(checkBytes[:len(signature)]) == string(signature) {
		fmt.Println("This is a valid PNG file.")
	} else {
		fmt.Println("This is not a valid PNG file.")
		fmt.Printf("Read bytes: %v\n", checkBytes[:n])
	}
}

// func main() {
// 	file, err := OpenPNGfile()
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
//
// 	ValidatePNG(file)
// }
