package main

import "fmt"

const englishHello = "Hello, "
const spanishHola = "Hola, "
const frenchBonjour = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchBonjour
	case "Spanish":
		prefix = spanishHola
	default:
		prefix = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("Mike", "English"))
}
