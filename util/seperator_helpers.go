package util

import (
	"encoding/hex"
	"log"
	"strings"
	"unicode/utf8"
)

func DetermineSeparator(delimiter string) rune {
	var separator rune

	if delimiter == "tab" {
		separator = '\t'
	} else if strings.Index(delimiter, "0x") == 0 {
		dlm, hex_err := hex.DecodeString(delimiter[2:])

		if hex_err != nil {
			log.Fatalln(hex_err)
		}

		separator, _ = utf8.DecodeRuneInString(string(dlm))
	} else {
		separator, _ = utf8.DecodeRuneInString(delimiter)
	}
	return separator
}
