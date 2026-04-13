// https://zap-hosting.com/guides/docs/dedicated-linux-dods/

// TODO:
// maybe add a GenericSource1 and just initialise it here

package dods

import (
	"fmt"
	"log"

	"github.com/kuche1/game-server/libgame"
	"github.com/kuche1/game-server/libsettings"
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
	if libutil.Exists(g.folder) {
		log.Fatalf("Path exists: %v\n", g.folder)
	}

	libsteamcmd.ServerInstall(g.folder, gameID)

	settings := _NewSettings()
	libsettings.Save(g.folder, settings)
}

func (g *GameDayOfDefeatSource) UpdateServer() {
	libsteamcmd.ServerUpdate(g.folder, gameID)
}

func (g *GameDayOfDefeatSource) StartServer(blocking bool) {
	settings := _NewSettings()
	libsettings.Load(g.folder, settings)

	libutil.Exec(
		"./srcds_run",
		true,
		g.folder,
		blocking,
		"-game", "dod",
		"-port", fmt.Sprintf("%v", settings.Port),
		"+map", settings.StartingMap, // mandatory (or at least for dods)
	)
}
