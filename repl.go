package main

import "fmt"
import "bufio"
import "os"
import "strings"


func startRepl(){

    for{
    fmt.Print("Pokedex >>")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()

    fmt.Println("echoing: ", text)
    }
}

func cleanInput(str string) []string {
    lowered := strings.ToLower(str)
    words := strings.Fields(lowered)

    return words
}
