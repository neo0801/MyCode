package main

import (
	"reflect"
	"fmt"
)

const prefixWordInEnglish = "Hello, "
const prefixWordInChinese = "你好, "
const prefixWordInSpanish = "Hola, "

// GetDefaultName returns default name according to language string
func GetDefaultName(language string) string {
	name := "World"
	switch language {
	case "chinese":
		name = "世界"
	case "spanish":
		name = "Elodie"
	}
	return name
}

// GetGreetingPrefix returns hello prefix according to language string
func GetGreetingPrefix(language string) string {
	prefix := prefixWordInEnglish
	switch language {
	case "chinese":
		prefix = prefixWordInChinese
	case "spanish":
		prefix = prefixWordInSpanish
	}
	return prefix
}

// Hello just return hello world!
func Hello(name, languange string) string {
	if name == "" {
		name = GetDefaultName(languange)
	}
	return GetGreetingPrefix(languange) + name
}

func main() {
	const pp = 3.1415
	fmt.Println(reflect.TypeOf(pp))
	fmt.Println(Hello("", ""))
}
