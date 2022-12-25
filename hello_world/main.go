package main

import (
	"fmt"
)

const spanish = "Spanish"
const franch = "Franch"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const franchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case franch:
		prefix = franchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("Chris", "Franch"))
}
