package module1

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
	"testing"
)

func TestModule1CheckEnv(t *testing.T) {
	actual := os.Getenv("GOPATH")
	expected0 := ""
	expected1 := path.Join(os.Getenv("HOME"), "go")

	if actual != expected0 && actual != expected1 {
		t.Errorf("environment variable GOPATH not set properly")
	}
}

func TestModule1CheckEnvGOOS(t *testing.T) {
	found := OpenFileAndFindNthString("module1.txt", 0, "GOOS")

	if !found {
		t.Errorf("'go env' does not work as expected")
	}
}

// OpenFileAndFindString opens a file and return if the given string is found or not
func OpenFileAndFindString(filename string, expected string) bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		trimmed := strings.Trim(t, " ")
		if trimmed == "" {
			continue
		}

		// matching logic
		if trimmed == expected {
			return true
		}
	}

	return false
}

// OpenFileAndFindNthString opens a file, look for Nth string splitted by a '=', and return if given expected string is found or not
func OpenFileAndFindNthString(filename string, nth int, expected string) bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		trimmed := strings.Trim(t, " ")
		if trimmed == "" {
			continue
		}

		// matching logic
		ss := strings.Split(trimmed, "=")
		if ss[nth] == expected {
			return true
		}
	}

	return false
}