package dods

import "github.com/kuche1/game-server/libgame"

type GameDayOfDefeatSource struct {
	gameFolder string
}

func NewGameDayOfDefeatSource(gameFolder string) libgame.Game {
	return &GameDayOfDefeatSource{
		gameFolder: gameFolder,
	}
}

func (g *GameDayOfDefeatSource) CreateServer() {
	panic("Not Implemented") // TODO
}

func (g *GameDayOfDefeatSource) UpdateServer() {
	panic("Not Implemented") // TODO
}

func (g *GameDayOfDefeatSource) StartServerForeground() {
	panic("Not Implemented") // TODO
}

func (g *GameDayOfDefeatSource) StartServerFork() {
	panic("Not Implemented") // TODO
}
