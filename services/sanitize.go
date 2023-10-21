package services

import (
	"regexp"

	iuliia "github.com/mehanizm/iuliia-go"
)

func Sanitize(str string) string {
	return regexp.MustCompile("[ ,.&*(!@#$~`\"'\\[\\]{}]").ReplaceAllString(iuliia.Wikipedia.Translate(str), "_")
}
