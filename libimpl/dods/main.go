// https://zap-hosting.com/guides/docs/dedicated-linux-dods/

package dods

import (
	"github.com/kuche1/game-server/libgame"
	"github.com/kuche1/game-server/libsteamcmd"
	"github.com/kuche1/game-server/libutil"
)

const gameID = 232290

type GameDayOfDefeatSource struct {
	folder string
}

func NewGameDayOfDefeatSource(gameFolder string) libgame.Game {
	return &GameDayOfDefeatSource{
		folder: gameFolder,
	}
}

func (g *GameDayOfDefeatSource) CreateServer() {
	libsteamcmd.ServerInstall(g.folder, gameID)
}

func (g *GameDayOfDefeatSource) UpdateServer() {
	libsteamcmd.ServerUpdate(g.folder, gameID)
}

func (g *GameDayOfDefeatSource) StartServer() {
	libutil.Exec(
		"./srcds_run",
		true,
		g.folder,
		"-game", "dod",
		"+map", "dod_anzio", // mandatory
	)
}

func (g *GameDayOfDefeatSource) StartServerFork() {
	panic("Not Implemented") // TODO
}
