package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenshHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenshHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main()  {
	fmt.Println(Hello("world", ""))
}