package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	gameNames := make([]string, 0, len(GameList))

	for name := range GameList {
		gameNames = append(gameNames, name)
	}

	var selectedGameName string

	prompt := &survey.Select{
		Message: "Select game:",
		Options: gameNames,
	}

	err := survey.AskOne(prompt, &selectedGameName)
	if err != nil {
		log.Fatal(err)
	}

	name := ""
	prompt := &survey.Input{
		Message: "Enter your name:",
	}

	survey.AskOne(prompt, &name)

	fmt.Printf("Hello, %s!\n", name)

	// game := gameNames[selectedGameName]()
}
