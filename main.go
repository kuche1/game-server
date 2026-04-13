package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	pGameName := flag.String("game", "", "Game name")
	gameFolder := flag.String("folder", "", "Game folder")
	action := flag.String("action", "", "Action to perform")

	flag.Parse()

	fmt.Printf("pGameName=%v\n", *pGameName)   // TODO: use
	fmt.Printf("gameFolder=%v\n", *gameFolder) // TODO: use
	fmt.Printf("action=%v\n", *action)         // TODO: use

	gameName := *pGameName

	if gameName == "" {
		fmt.Fprintf(os.Stderr, "Game not specified\n")
		flag.Usage()
		os.Exit(1)
	}

	gameNames := getGameNames()
	if !slices.Contains(gameNames, gameName) {
		fmt.Fprintf(os.Stderr, "Invalid game:\n")
		fmt.Fprintf(os.Stderr, "- %v\n", gameName)
		fmt.Fprintf(os.Stderr, "Valid games:\n")
		for _, name := range gameNames {
			fmt.Fprintf(os.Stderr, "- %v\n", name)
		}
		os.Exit(1)
	}
}
