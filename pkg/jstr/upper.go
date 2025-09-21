package jstr

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Function to convert string to title case (e.g. "hello world" -> "Hello World")
func Title(s string) string {
	cases := cases.Title(language.Und)
	return cases.String(s)
}

// Function to convert string to sentence case (e.g. "hello world" -> "Hello world")
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Function to convert snake_case to class camelCase (e.g. "hello_world" -> "HelloWorld")
func ClassCamel(s string) string {
	var result string
	comp := strings.Split(s, "_")
	for _, v := range comp {
		result += Title(v)
	}
	return result
}

// Function to convert snake_case to variable camelCase (e.g. "hello_world" -> "helloWorld")
func VariableCamel(s string) string {
	var result string
	comp := strings.Split(s, "_")
	for i, v := range comp {
		if i == 0 {
			result += v
		} else {
			result += Title(v)
		}
	}
	return result
}
