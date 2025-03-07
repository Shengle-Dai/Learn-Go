package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("bro", ""))
}
