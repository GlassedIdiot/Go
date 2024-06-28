package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"os"
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

//My idea for this to validate. Read only the first 4bytes using a for loop append to a byte slice and then use the docs example to validate. Will do it later.  

func (head *Header)ValidatePNG( breader *bytes.Reader)  {
  
  if err := binary.Read(breader,binary.BigEndian,head){}

}
