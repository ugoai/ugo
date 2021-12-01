package main

import (
	"bufio"
	"fmt"
	"os"

	core "ugo/core"
)

func main() {
	fmt.Print("Welcome to Ugo !\n\n")
	fmt.Println("Initialize my brain")
	ugoCore := core.Brain{
		Lang: "fr",
	}
	ugoCore.Init()
	fmt.Println("Initialize my webapp")
	fmt.Println("Ready on http://127.0.0.1:1010")
	reader := bufio.NewReader(os.Stdin)
	var phrase string
	fmt.Println("Enter text:") // Question to the user
	phrase, readerError := reader.ReadString('\n')
	if readerError != nil {
		return
	}
	ugoCore.AskUgo(phrase)
}
