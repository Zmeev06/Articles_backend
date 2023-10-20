package services

import iuliia "github.com/mehanizm/iuliia-go"

func Sanitize(str string) string {
	return iuliia.Wikipedia.Translate(str)
}
