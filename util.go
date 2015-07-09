package main

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

func genericReader(fileName string) (*bufio.Reader, error) {
	var reader io.Reader
	var err error = nil
	var f *os.File = nil

	if fileName == "" {
		log.Println("Reading from stdin")
		return bufio.NewReader(os.Stdin), nil
	} else {
		f, err = os.Open(fileName)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if strings.HasSuffix(fileName, ".gz") {
			reader, err = gzip.NewReader(f)
			if err != nil {
				log.Println(err)
				return nil, err
			}

		} else {
			if strings.HasSuffix(fileName, ".bz2") {
				reader = bzip2.NewReader(f)
			} else {
				reader = f
			}
		}

		log.Println("Opening: " + fileName)
		return bufio.NewReader(reader), nil
	}

}
