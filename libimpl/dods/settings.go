package dods

type _Settings struct {
	Port        uint32
	StartingMap string
}

func _NewSettings() *_Settings {
	return &_Settings{
		Port:        27015,
		StartingMap: "dod_anzio",
	}
}
