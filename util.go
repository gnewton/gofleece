package main

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const GZIP_SUFFIX = ".gz"
const BZIP2_SUFFIX = ".bz2"

var ErrZeroLengthFile = errors.New("File has zero length")

func genericReader(fileName string) (*bufio.Reader, error) {
	var reader io.Reader
	var err error = nil
	var f *os.File = nil

	if fileName == "" {
		if VERBOSE {
			log.Println("Reading from stdin")
		}
		return bufio.NewReader(os.Stdin), nil
	} else {
		if VERBOSE {
			log.Println("Opening: " + fileName)
		}
		f, err = os.Open(fileName)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		fileInfo, err := f.Stat()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if fileInfo.Size() == 0 {
			return nil, ErrZeroLengthFile
		}

		if VERBOSE {
			log.Println(" File size: " + strconv.FormatInt(fileInfo.Size(), 10))
		}

		if strings.HasSuffix(fileName, GZIP_SUFFIX) {
			reader, err = gzip.NewReader(f)
			if err != nil {
				log.Println(err)
				return nil, err
			}

		} else {
			if strings.HasSuffix(fileName, BZIP2_SUFFIX) {
				reader = bzip2.NewReader(f)
			} else {
				reader = f
			}
		}

		return bufio.NewReader(reader), nil
	}

}
