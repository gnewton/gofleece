package main

import (
	"encoding/json"
	"os"

	//"fmt"
	"bufio"
	//	"errors"

	"io"
	"log"
	"strings"
)

func verifyJson(reader *bufio.Reader) error {

	prefix, err := reader.Peek(128)
	if err != nil {
		if err != bufio.ErrBufferFull && err != io.EOF {
			logg(err)
			return err
		}

	}

	return parseJson(prefix, reader)
}

func parseJson(prefix []byte, jsonReader *bufio.Reader) error {
	if jsonHasLeadingSquareBrace(string(prefix)) {
		if VERBOSE {
			logg("Array")
		}
		return decodeJsonArray(jsonReader)
	} else {
		if VERBOSE {
			logg("Map")
		}
		return decodeJsonMap(jsonReader)
	}
}

func isValidJson(r *bufio.Reader) error {
	prefix, err := r.Peek(1)
	if err == io.EOF {
		return err
	}
	prefix, err = r.Peek(128)
	if err != nil {
		if err != bufio.ErrBufferFull && err != io.EOF {
			logg(err)
			return err
		}

	}
	if jsonHasLeadingSquareBrace(string(prefix)) {
		err = decodeJsonArray(r)
	} else {
		err = decodeJsonMap(r)
	}
	if err == io.EOF {
		return nil
	} else {
		return err
	}
}

func decodeJsonMap(r *bufio.Reader) error {
	var mp map[string]interface{}
	mp = nil

	dec := json.NewDecoder(r)

	for {
		if err := dec.Decode(&mp); err == io.EOF {
			break
		} else if err != nil {
			if err == io.EOF {
				return nil
			} else {

				return err
			}
		}
	}
	return nil
}

func decodeJsonArray(r *bufio.Reader) error {
	var mp []map[string]interface{}
	mp = nil

	dec := json.NewDecoder(r)

	for {
		if err := dec.Decode(&mp); err == io.EOF {
			break
		} else if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
	}
	return nil
}

func jsonHasLeadingSquareBrace(s string) bool {
	var trimmed string
	trimmed = strings.TrimSpace(s)
	if trimmed[0] == '[' {
		return true
	}
	return false
}

func fatal(err error) {
	if SILENT || err == nil {
		os.Exit(1)
	} else {
		log.Fatal(err)
	}
}

func logg(s ...interface{}) {
	if !SILENT && s != nil && len(s) > 0 {
		log.Println(s[0])
	}
}

func logv(s ...interface{}) {
	if VERBOSE {
		logg(s)
	}
}
