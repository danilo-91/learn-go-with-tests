package main

import "fmt"

const hola = "Hola"
const mundo = "Mundo"
const bonjour = "Bonjour"
const monde = "Monde"

func Hello(name, lang string) string {
    hello := "Hello"
    switch lang {
    case "es":
        hello = "Hola"
    case "fr":
        hello = "Bonjour"
    case "it":
        hello = "Ciao"
    }

    if name == "" {
        name = world(lang)
    }

    return fmt.Sprintf("%v, %v!", hello, name) 
}

func world(lang string) string {
    switch lang {
    case "es":
        return "Mundo"
    case "fr":
        return "Monde"
    case "it":
        return "Mondo"
    default:
        return "World"
    }
}

        

func main() {
    fmt.Println(Hello("Danilo", "en"))
}

