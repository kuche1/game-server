package main

import (
	"github.com/kuche1/game-server/libgame"
	"github.com/kuche1/game-server/libimpl/dods"
)

var GameList = map[string](func(gameFolder string) libgame.Game){
	"Day of Defeat: Source": dods.NewGameDayOfDefeatSource,
}

func getGameNames() []string {
	names := make([]string, 0, len(GameList))

	for name := range GameList {
		names = append(names, name)
	}

	return names
}
