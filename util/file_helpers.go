package util

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func OpenFileOrStdin(path *string) *os.File {
	var fp *os.File
	var err error

	if (*path) == "stdin" {
		fp = os.Stdin
		err = nil
	} else {
		fp, err = os.Open(*cleanPath(path))
	}

	if err != nil {
		log.Fatalln(err)
	}

	return fp
}

func cleanPath(path *string) *string {
	var result string
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	if (*path)[:2] == "~/" {
		dir := usr.HomeDir + "/"
		result = strings.Replace(*path, "~/", dir, 1)
	} else {
		result = (*path)
	}

	abs_result, abs_err := filepath.Abs(result)
	if abs_err != nil {
		log.Fatalln(err)
	}

	clean_result := filepath.Clean(abs_result)

	return &clean_result
}

func RewindFile(fileHandle *os.File) {
	_, rewind_err := fileHandle.Seek(0, 0)

	if rewind_err != nil {
		log.Fatalln("Unable to rewind file")
	}
}
