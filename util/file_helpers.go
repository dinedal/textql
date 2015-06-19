package util

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func OpenFileOrStdDev(path string) *os.File {
	var fp *os.File
	var err error

	if path == "stdin" {
		fp = os.Stdin
		err = nil
	} else if path == "stdout" {
		fp = os.Stdout
		err = nil
	} else {
		fp, err = os.Open(CleanPath(path))
	}

	if err != nil {
		log.Fatalln(err)
	}

	return fp
}

func CleanPath(path string) string {
	result := ""

	if path == "" {
		return ""
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	if path[:2] == "~/" {
		dir := usr.HomeDir + "/"
		result = strings.Replace(path, "~/", dir, 1)
	} else {
		result = path
	}

	absResult, absErr := filepath.Abs(result)
	if absErr != nil {
		log.Fatalln(absErr)
	}

	cleanResult := filepath.Clean(absResult)

	return cleanResult
}

func RewindFile(fileHandle *os.File) {
	_, rewindErr := fileHandle.Seek(0, 0)

	if rewindErr != nil {
		log.Fatalln("Unable to rewind file")
	}
}
