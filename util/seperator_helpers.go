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
		dlm, hexErr := hex.DecodeString(delimiter[2:])

		if hexErr != nil {
			log.Fatalln(hexErr)
		}

		separator, _ = utf8.DecodeRuneInString(string(dlm))
	} else {
		separator, _ = utf8.DecodeRuneInString(delimiter)
	}
	return separator
}
