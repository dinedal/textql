package storage

import "regexp"

func regExp(re, s string) (bool, error) {
	return regexp.MatchString(re, s)
}
