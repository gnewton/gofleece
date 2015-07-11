// from https://www.socketloop.com/tutorials/golang-decode-unmarshal-unknown-json-data-type-with-map-string-interface
package main

import (
	"encoding/json"
	"os"
	"runtime"
	//"fmt"
	"bufio"
	//	"errors"
	"flag"
	"io"
	"log"
	"strings"
)

func handleParameters() error {
	flag.Parse()
	return nil
}

var STOP_ON_ERROR bool = false
var ZERO_LENGTH_FILE_IS_VALID bool = false
var VERBOSE bool = false
var SILENT bool = false

func init() {
	flag.BoolVar(&STOP_ON_ERROR, "e", STOP_ON_ERROR, "Stop if error encountered")
	flag.BoolVar(&SILENT, "s", SILENT, "Do not produce any output")
	flag.BoolVar(&VERBOSE, "v", VERBOSE, "Verbosity of logging")
	flag.BoolVar(&ZERO_LENGTH_FILE_IS_VALID, "z", ZERO_LENGTH_FILE_IS_VALID, "Acccept a file of zero length as being valid JSON")
}

func main() {
	var finalErr error = nil
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var fileNames = []string{""}
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := handleParameters()
	if err != nil {
		flag.Usage()
		fatal(err)
	}

	if len(flag.Args()) != 0 {
		fileNames = flag.Args()
	}

	for _, fileName := range fileNames {
		err = verifyJson(fileName)
		if err != nil {
			logg("###ERROR \"" + fileName + "\": NOT valid JSON")
			logg(err)
			finalErr = err
			if STOP_ON_ERROR {
				fatal(err)
			}
		} else {
			logv("Valid JSON")
		}
		logv("------------------")
	}
	if finalErr != nil {
		fatal(nil)
	}
}

func verifyJson(fileName string) error {
	jsonStream, err := genericReader(fileName)
	if err != nil {
		if err == ErrZeroLengthFile && ZERO_LENGTH_FILE_IS_VALID {
			return nil
		}
		return err
	}

	prefix, err := jsonStream.Peek(128)
	if err != nil {
		if err != bufio.ErrBufferFull && err != io.EOF {
			logg(err)
			return err
		}

	}

	return parseJson(prefix, jsonStream)
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
