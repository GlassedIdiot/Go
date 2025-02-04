package utils

import (
	"bytes"
	"fmt"
	"hats/Stegnography/Models"
	"io"
	"log"
	"os"
	"strconv"
)

func WritePayload(r *bytes.Reader, c *models.CmdLineOpts, b []byte) {
	offset, err := strconv.ParseInt(c.Offset, 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse offset: %s", err)
	}

	w, err := os.Create(c.Output)
	if err != nil {
		log.Fatalf("File was not created: %s", err)
	}
	defer w.Close()

	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("Seek failed: %s", err)
	}

	buff := make([]byte, offset)
	n, err := r.Read(buff)
	if err != nil && err != io.EOF {
		log.Fatalf("Read failed: %s", err)
	} else if n == 0 {
		log.Fatal("No bytes read")
	}

	_, err = w.Write(buff[:n])
	if err != nil {
		log.Fatalf("Write failed: %s", err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Fatalf("Write failed: %s", err)
	}

	_, err = io.Copy(w, r)
	if err != nil {
		log.Fatalf("Copy failed: %s", err)
	}

	fmt.Printf("Success: %s created\n", c.Output)
}

