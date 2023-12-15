package main

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func returnHello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name + "!"
	} else if language == "French" {
		return frenchHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}
