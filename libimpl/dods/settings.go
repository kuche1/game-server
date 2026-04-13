package dods

type _Settings struct {
	StartingMap string
}

func _NewSettings() *_Settings {
	return &_Settings{
		StartingMap: "dod_anzio",
	}
}
