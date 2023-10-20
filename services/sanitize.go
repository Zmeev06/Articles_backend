package services

import (
	"strings"

	iuliia "github.com/mehanizm/iuliia-go"
)

func Sanitize(str string) string {
	return strings.ReplaceAll(iuliia.Wikipedia.Translate(str), " ", "_")
}
