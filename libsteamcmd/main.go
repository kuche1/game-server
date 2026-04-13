package libsteamcmd

import (
	"fmt"
	"log"

	"github.com/kuche1/game-server/libutil"
)

func ServerInstall(forceInstallDir string, gameID int64) {
	if libutil.Exists(forceInstallDir) {
		log.Fatalf("Path exists: %v\n", forceInstallDir)
	}

	runCustom(forceInstallDir, "+app_update", fmt.Sprintf("%v", gameID), "validate")
}

func ServerUpdate(forceInstallDir string, gameID int64) {
	if !libutil.IsFolder(forceInstallDir) {
		log.Fatalf("Not a folder: %v\n", forceInstallDir)
	}

	runCustom(forceInstallDir, "+app_update", fmt.Sprintf("%v", gameID))
}
