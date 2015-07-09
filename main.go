// from https://www.socketloop.com/tutorials/golang-decode-unmarshal-unknown-json-data-type-with-map-string-interface
package main

import (
	"encoding/json"
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

var DEBUG bool

func init() {
	flag.BoolVar(&DEBUG, "d", DEBUG, "Debug; prints out much information")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := handleParameters()
	if err != nil {
		flag.Usage()
		log.Fatal(err)
	}

	var fileName string
	if len(flag.Args()) > 0 {
		fileName = flag.Args()[0]
	}

	//fileName := "/run/media/gnewton/NewtonNTFS/data/json/metadata/repository_metadata_2014-06-06_150.json"
	//fileName := "repository_metadata_2014-06-06_150.json.bz2"
	//f, err := os.Open("/home/gnewton/Downloads/repository_metadata_2013-12-12_127.json")
	//f, err := os.Open("/run/media/gnewton/NewtonNTFS/data/json/metadata/repository_metadata_2013-12-16_143.json")
	//f, err := os.Open("/tmp/a.json")
	//msg := "{\"assets\" : {\"old\" : 123}, \"foo\":[1,2,3], \"person\":{\"age\":27, \"sex\":\"m\"}, \"smith\": 56}"

	jsonStream, err := genericReader(fileName)
	if err != nil {
		log.Fatal(err)
	}

	prefix, err := jsonStream.Peek(128)
	if err != nil {
		if err != bufio.ErrBufferFull {
			//log.Fatal(err)
		}
	}
	if jsonHasLeadingSquareBrace(string(prefix)) {
		err = isValidArray(jsonStream)
	} else {
		err = isValidMap(jsonStream)
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
			log.Println(err)
		}
	}
	if jsonHasLeadingSquareBrace(string(prefix)) {
		err = isValidArray(r)
	} else {
		err = isValidMap(r)
	}
	if err == io.EOF {
		err = nil
	}
	return err
}

func isValidMap(r *bufio.Reader) error {
	var mp map[string]interface{}
	mp = nil

	dec := json.NewDecoder(r)
	dec.UseNumber()

	for {
		if err := dec.Decode(&mp); err == io.EOF {
			break
		} else if err != nil {
			if err == io.EOF {
				return nil
			} else {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

func isValidArray(r *bufio.Reader) error {
	var mp []map[string]interface{}
	mp = nil

	dec := json.NewDecoder(r)
	dec.UseNumber()

	for {
		if err := dec.Decode(&mp); err == io.EOF {
			break
		} else if err != nil {
			if err == io.EOF {
				return nil
			} else {
				log.Println(err)
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
