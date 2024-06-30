package steganography

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"hats/Stegnography/Models"
	"log"
	"os"
	"strconv"
)

type Header struct {
	Header uint64
}
type Chunk struct {
	size uint32
	Type uint32
	Data []byte
	CRC  uint32
}
type MetaChunk struct {
	Chk    Chunk
	Offset int64
}

func ErrorCheck(e error) {
	log.Fatal(e)
}

func PreProcessing(file *os.File) (*bytes.Reader, error) {
	// Getting the stats of the file
	stats, err := file.Stat()
	if err != nil {
		os.Exit(1)
		return nil, err

	}

	// Getting the size of the file
	size := stats.Size()
	b := make([]byte, size)

	// Created a new rread for reading
	someReader := bufio.NewReader(file)

	_, err = someReader.Read(b)
	if err != nil {
		log.Fatalf("%s", err)
		return nil, err
	}

	bReader := bytes.NewReader(b)

	return bReader, nil
}

// My idea for this to validate. Read only the first 4bytes using a for loop append to a byte slice and then use the docs example to validate. Will do it later.
// [x]My idea for this to validate. Read only the first 4bytes using a for loop append to a byte slice and then use the docs example to validate. Will do it later.Did this works well.

func (head *Header) ValidatePNG(breader *bytes.Reader) {
	if err := binary.Read(breader, binary.BigEndian, &head.Header); err != nil {
		log.Fatal(err)
	}

	Bcheck := make([]byte, 8)

	binary.BigEndian.PutUint64(Bcheck, head.Header)

	if string(Bcheck[1:4]) == "PNG" {
		fmt.Printf("%s:%s\n", "DetectedFileType", "PNG")
	} else {
		fmt.Print("Not a PNG file.")
		return
	}
}

func (mc *MetaChunk) ParsePNG(b *bytes.Reader, c *models.CmdLineOpts) {
	// Skipping some things and going for the main block right now.

	count := 1
	chunkType := ""
	endChunkType := "IEND"

	for chunkType != endChunkType {

		fmt.Println("---- Chunk # " + strconv.Itoa(count) + " ----")

		fmt.Printf("Chunks-Offset=%d\n", mc.Offset)
		fmt.Printf("Length=%d\n", mc.Chk.size)
		fmt.Printf("Type of data:%s\n", mc.Chk.Data)
		fmt.Printf("Chunk CRC = %x\n", mc.Chk.CRC)

	}
}

func (mc *MetaChunk) FindingOffset(b *bytes.Reader) {
	offset, err := b.Seek(0, 1)

	checkErr(err)

	mc.Offset = offset
}
