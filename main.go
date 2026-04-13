package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/kuche1/game-server/libgame"
)

func main() {
	pGameName := flag.String("game", "", "Game name")
	pGameFolder := flag.String("folder", "", "Game folder")
	pAction := flag.String("action", "", "Action to perform")

	flag.Parse()

	///// game

	gameName := *pGameName

	if gameName == "" {
		fmt.Fprintf(os.Stderr, "Game not specified, see -help\n")
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

	///// folder

	gameFolder := *pGameFolder

	if gameFolder == "" {
		fmt.Fprintf(os.Stderr, "Folder not specified, see -help\n")
		os.Exit(1)
	}

	///// action

	action := *pAction

	if action == "" {
		fmt.Fprintf(os.Stderr, "Folder not specified, see -help\n")
		os.Exit(1)
	}

	///// run

	gameImpl := GameList[gameName](gameFolder)

	badActionMessage := libgame.RunAction(gameImpl, action)
	if badActionMessage != "" {
		fmt.Fprint(os.Stderr, badActionMessage)
		os.Exit(1)
	}
}
