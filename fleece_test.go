package main

import (
	"bufio"
	"log"
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

func TestValidArray1(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("[" + validJsonMap0 + "]"))
	err := isValidJson(reader)
	if err != nil {
		t.FailNow()
	}
}

func TestInvalidArray0(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("z" + validJsonArray0))
	err := isValidJson(reader)
	log.Println(err)
	if err == nil {
		t.FailNow()
	}
}

func TestInvalidArray1(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("[" + validJsonArray0 + "]"))
	err := isValidJson(reader)
	log.Println(err)
	if err == nil {
		t.FailNow()
	}
}

func TestValidMap(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(validJsonMap0))
	err := isValidJson(reader)
	if err != nil {
		t.FailNow()
	}
}

func TestInvalidMap0(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("z" + validJsonMap0))
	err := isValidJson(reader)
	if err == nil {
		t.FailNow()
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
