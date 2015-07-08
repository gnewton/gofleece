package main

import (
	"bufio"

	"strings"
	"testing"
)

func TestValidArray0(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(validJsonArray0))
	err := isValidJson(reader)
	if err != nil {
		t.FailNow()
	}
}

func TestValidArrays(t *testing.T) {
	for i, _ := range validArrays {
		reader := bufio.NewReader(strings.NewReader(validArrays[i]))
		err := isValidJson(reader)
		if err != nil {
			t.FailNow()
		}
	}
}

func TestValidMaps(t *testing.T) {
	for i, _ := range validMaps {
		reader := bufio.NewReader(strings.NewReader(validMaps[i]))
		err := isValidJson(reader)
		if err != nil {
			t.FailNow()
		}
	}
}

func TestValidMapsAsWrappedAsArrays(t *testing.T) {
	for i, _ := range validMaps {
		reader := bufio.NewReader(strings.NewReader("[" + validMaps[i] + "]"))
		err := isValidJson(reader)
		if err != nil {
			t.FailNow()
		}
	}
}

func TestValidInvalidArrays(t *testing.T) {
	for i, _ := range validArrays {
		reader := bufio.NewReader(strings.NewReader("z HH " + validArrays[i]))
		err := isValidJson(reader)
		if err == nil {
			t.FailNow()
		}
	}
}

func TestInvalidMaps(t *testing.T) {
	for i, _ := range validMaps {
		reader := bufio.NewReader(strings.NewReader("mmmm " + validMaps[i]))
		err := isValidJson(reader)
		if err == nil {
			t.FailNow()
		}
	}

}

func TestEmptyString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	err := isValidJson(reader)
	if err == nil {
		t.FailNow()
	}
}

func TestSmallJson(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("{}"))
	err := isValidJson(reader)
	if err != nil {
		t.FailNow()
	}
}
