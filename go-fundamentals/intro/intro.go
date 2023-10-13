package main

import (
	"fmt"
)

const (
	french = "French"
	spanish = "Spanish"
	japanese = "Japanese"

	 englishHelloPrefix = "Hello, "
	 spanishHelloPrefix = "Hola, "
	 frenchHelloPrefix = "Bounjour, "
	 japaneseHelloPrefix = "Konichiwa, "

)


func Hello(name string, lang string) string {
	
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		case japanese:
			prefix = japaneseHelloPrefix
		default:
			prefix = englishHelloPrefix
		}
		return 
}

func main(){
	fmt.Println(Hello("",""))
}

