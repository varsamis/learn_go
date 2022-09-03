package main

import "fmt"

const french = "French"
const greek = "Greek"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const greekHelloPrefix = "Γειά σου, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Holla, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case greek:
		prefix = greekHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
