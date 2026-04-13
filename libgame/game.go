package libgame

type Game interface {
	CreateServer()
	UpdateServer()
	StartServerForeground()
	StartServerFork()
}
