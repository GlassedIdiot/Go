package Steganography

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	models "hats/Stegnography/Models"
	"hats/Stegnography/utils"
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

func (mc *MetaChunk) ValidatePNG(breader *bytes.Reader) {
	var head Header
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

		mc.ReadChunk(b)
		fmt.Printf("Chunks-Offset=%d\n", mc.Offset)
		fmt.Printf("Length=%d\n", mc.Chk.size)
		fmt.Printf("Type of data:%d\n", mc.Chk.Type)
		fmt.Printf("Chunk CRC = %x\n", mc.Chk.CRC)

		count++
		// Gonna create some helper functions to read each Chunk Type.
	}
}

func (mc *MetaChunk) ProcessImage_Payload(b *bytes.Reader, c *models.CmdLineOpts) {
	// Now how shall we write the pay load will do the encoding part later.
	// -[x] Adding simple payload delivery.
	//-[ ] Add the encoding part for writing the payload.

	var m MetaChunk

	m.Chk.size = mc.CreateChunkSize()

	Temp, err := strconv.ParseUint(c.Type, 10, 32)
	if err != nil {
		ErrorCheck(err)
	}
	m.Chk.Type = uint32(Temp)

	m.Chk.Data = []byte(c.Payload)

	m.Chk.CRC = mc.CreateCRCcheck()

	bm := mc.marshalData()
	bmb := bm.Bytes()
	fmt.Printf("Payload Original: % X\n", []byte(c.Payload))
	fmt.Printf("Payload: % X\n", m.Chk.Data)
	utils.WritePayload(b, c, bmb)
}

func (mc *MetaChunk) ReadChunk(b *bytes.Reader) {
	mc.ReadSize(b)
	mc.ReadType(b)
	mc.ReadChunkBytes(b)
	mc.ReadCRC(b)
}

func (mc *MetaChunk) ReadSize(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.size); err != nil {
		ErrorCheck(err)
	}
}

func (mc *MetaChunk) ReadType(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Type); err != nil {
		ErrorCheck(err)
	}
}

func (mc *MetaChunk) ReadChunkBytes(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.Data); err != nil {
		ErrorCheck(err)
	}
}

func (mc *MetaChunk) ReadCRC(b *bytes.Reader) {
	if err := binary.Read(b, binary.BigEndian, &mc.Chk.CRC); err != nil {
		ErrorCheck(err)
	}
}

func (mc *MetaChunk) FindingOffset(b *bytes.Reader) {
	offset, err := b.Seek(0, 1)

	checkErr(err)

	mc.Offset = offset
}

func (mc *MetaChunk) CreateChunkSize() uint32 {
	return uint32(len(mc.Chk.Data))
}

func (mc *MetaChunk) CreateCRCcheck() uint32 {
	bytesLSB := new(bytes.Buffer)

	if err := binary.Write(bytesLSB, binary.BigEndian, mc.Chk.Type); err != nil {
		ErrorCheck(err)
	}

	if err := binary.Write(bytesLSB, binary.BigEndian, mc.Chk.Data); err != nil {
		ErrorCheck(err)
	}

	return crc32.ChecksumIEEE(bytesLSB.Bytes())
}

func (mc *MetaChunk) marshalData() *bytes.Buffer {
	bytesMSB := new(bytes.Buffer)
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.size); err != nil {
		log.Fatal(err)
	}
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Type); err != nil {
		log.Fatal(err)
	}
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.Data); err != nil {
		log.Fatal(err)
	}
	if err := binary.Write(bytesMSB, binary.BigEndian, mc.Chk.CRC); err != nil {
		log.Fatal(err)
	}

	return bytesMSB
}
