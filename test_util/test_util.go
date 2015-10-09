package test_util

import (
	"io/ioutil"
	"os"
)

func OpenFileFromString(contents string) *os.File {
	f, _ := ioutil.TempFile("./", "csv")
	f.WriteString(contents)
	f.Seek(0, 0)
	return f
}
