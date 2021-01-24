// gofleece
// Copyright (C) 2015  Glen Newton

// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package main

import (
	"flag"
	"log"
)

func handleParameters() error {
	flag.Parse()
	return nil
}

var STOP_ON_ERROR bool = false
var ZERO_LENGTH_FILE_IS_VALID bool = false
var VERBOSE bool = false
var SILENT bool = false
var PROCESS_XML bool = false

func init() {
	flag.BoolVar(&STOP_ON_ERROR, "e", STOP_ON_ERROR, "Stop if error encountered")
	flag.BoolVar(&SILENT, "s", SILENT, "Do not produce any output")
	flag.BoolVar(&VERBOSE, "v", VERBOSE, "Verbosity of logging")
	flag.BoolVar(&ZERO_LENGTH_FILE_IS_VALID, "z", ZERO_LENGTH_FILE_IS_VALID, "Acccept a file of zero length as being valid JSON")
	flag.BoolVar(&PROCESS_XML, "x", PROCESS_XML, "Process XML instead of JSON (default false)")
}

func main() {
	var finalErr error = nil
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var fileNames = []string{""}

	err := handleParameters()
	if err != nil {
		flag.Usage()
		fatal(err)
	}

	if len(flag.Args()) != 0 {
		fileNames = flag.Args()
	}

	for _, filename := range fileNames {
		err = process(filename)

		if err != nil {
			logg("###ERROR \"" + filename + "\": NOT valid JSON")
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

func process(filename string) error {
	reader, err := genericReader(filename)
	if err != nil {
		if err == ErrZeroLengthFile && ZERO_LENGTH_FILE_IS_VALID {
			return nil
		}
		return err
	}

	if PROCESS_XML {
		return nil
	} else {
		return verifyJson(reader)
	}
}
