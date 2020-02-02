package main

import "fmt"

const prefixWordInEnglish = "Hello, "
const prefixWordInChinese = "你好, "
const prefixWordInSpanish = "Hola, "

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
	fmt.Println(Hello("", ""))
}
