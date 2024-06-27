package main

import (
	"bufio"
	"bytes"
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
	stats, err := file.Stat()
	if err != nil {
		os.Exit(1)
		return nil, err

	}

	size := stats.Size()
	b := make([]byte, size)

	someReader := bufio.NewReader(file)

	_, err = someReader.Read(b)
	if err != nil {
		log.Fatalf("%s", err)
		return nil, err
	}

	bReader := bytes.NewReader(b)

	return bReader, nil
}
