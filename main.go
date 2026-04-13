package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	fmt.Printf("hi\n")

	var choice int

	prompt := &survey.Select{
		Message: "Choose an option:",
		Options: []string{"Start", "Settings", "Quit"},
	}

	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("choice=%v\n", choice)
}
