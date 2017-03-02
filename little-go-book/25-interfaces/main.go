package main

import (
	"fmt"
)

// Logger interface - Something that does some logging.
// This is an interface similar to how they're defined in other languages. Just
// a contract of funcionality, no actual function bodies.
type Logger interface {
	Log(message string)
}

// Implementing an interface in a structure is not explicit like some languages
// (e.g. class AThing implements ThingInterface {}). Instead the implementation
// is implicit if a structure has a function with a matching name, parameters
// and return type. In this case we need a 'Log' function that takes a message
// string parameter and returns nothing.

// BoringLogger - Does some logging. A bit boring.
type BoringLogger struct{}

// Log (BoringLogger) - Logs a message in a boring way.
func (b BoringLogger) Log(message string) {
	fmt.Printf("New log: %s\n", message)
}

// FancyLogger - Does some logging, pretty fancy tbh.
type FancyLogger struct{}

// Log (FancyLogger) - Logs a message in a fancy way.
func (b FancyLogger) Log(message string) {
	fmt.Println("******")
	fmt.Printf("%s!\n", message)
	fmt.Println("******")
}

func main() {
	// Create the loggers.
	aBoringLogger := BoringLogger{}
	aFancyLogger := FancyLogger{}

	// Process the messages.
	process(aBoringLogger, "Something happened")
	process(aFancyLogger, "Something happened")
}

// Process a message and log it with the specified logger.
func process(logger Logger, message string) {
	logger.Log(message)
}
