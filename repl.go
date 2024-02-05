package main

import "fmt"
import "bufio"
import "os"


func startRepl(){

    for{
    fmt.Print("Pokedex >>")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()

    fmt.Println("echoing: ", text)
}

}
