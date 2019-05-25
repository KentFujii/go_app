package service

import (
	"strings"
	"errors"
)

type String struct{}

func (String) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return strings.ToUpper(s), nil
}

func (String) Count(s string) int {
	return len(s)
}
