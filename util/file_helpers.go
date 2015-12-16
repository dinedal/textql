package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func IsPathDir(path string) bool {
	fp, err := os.Open(CleanPath(path))

	if err != nil {
		log.Fatalln(err)
	}

	defer fp.Close()

	stat, statErr := fp.Stat()

	if statErr != nil {
		log.Fatalln(statErr)
	}

	return stat.IsDir()
}

func OpenFileOrStdDev(path string, write bool) *os.File {
	var fp *os.File
	var err error

	if path == "stdin" {
		fp = os.Stdin
		err = nil
	} else if path == "stdout" {
		fp = os.Stdout
		err = nil
	} else {
		if write {
			fp, err = os.Create(CleanPath(path))
		} else {
			fp, err = os.Open(CleanPath(path))
		}
	}

	if err != nil {
		log.Fatalln(err)
	}

	stat, statErr := fp.Stat()

	if statErr != nil {
		log.Fatalln(err)
	}

	if stat.IsDir() {
		log.Fatalf("%s: is a directory\n", path)
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

	if len(path) > 1 && path[:2] == "~/" {
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

func IsThereDataOnStdin() bool {
	stat, statErr := os.Stdin.Stat()

	if statErr != nil {
		log.Fatalln(statErr)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	} else {
		return false
	}
}

func AllFilesInDirectory(path string) []string {
	cleanPath := CleanPath(path)
	directoryEntries, err := ioutil.ReadDir(cleanPath)
	result := make([]string, 0)

	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range directoryEntries {
		if !entry.IsDir() {
			result = append(result, filepath.Join(cleanPath, entry.Name()))
		}
	}

	return result
}
