package nameutil

import (
	"regexp"
	"strings"
)

// ToExportedCamel conv name to exported-name. eg: resource -> Resource
func ToExportedCamel(name string) string {
	if len(name) == 0 {
		panic("invalid name")
	}
	return strings.Title(name)
}

//ToShortName ...
func ToShortName(name string) string {
	if len(name) == 0 {
		panic("invalid name")
	}
	return name[0:1]
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

//ToSnakeCaseName ...
func ToSnakeCaseName(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

//ToUnexportedCamel ...
func ToUnexportedCamel(name string) string {
	return name
}
