package helper

import (
	"fmt"
	"testing"
)

func TestHelloInsun(t *testing.T) {
	result := SayHello("insun")
	fmt.Println(result)

	if result != "Hello insun" {
		panic("Result is not Hello insun")
	}
}

func TestHelloInzun(t *testing.T) {
	result := SayHello("inzun")
	if result != "Hello inzun" {
		panic("Result is not Hello insun")
	}
}
