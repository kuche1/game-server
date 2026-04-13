package libsettings

import (
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/kuche1/game-server/libgame"
	"github.com/kuche1/game-server/libutil"
)

func Save[T any](gameFolder string, settingsStruct *T) {
	filePath := path.Join(gameFolder, libgame.SettingsFileName)

	fileObj, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileObj.Close()

	if err := toml.NewEncoder(fileObj).Encode(settingsStruct); err != nil {
		log.Fatal(err)
	}
}

// Will create the settings file if it does not exist
func Load[T any](gameFolder string, settingsStruct *T) {
	filePath := path.Join(gameFolder, libgame.SettingsFileName)

	if !libutil.Exists(filePath) {
		log.Print("Settings file has been deleted, creating new one")
		Save(gameFolder, settingsStruct)
		return
	}

	_, err := toml.DecodeFile(filePath, settingsStruct)
	if err != nil {
		log.Fatal(err)
	}
}
