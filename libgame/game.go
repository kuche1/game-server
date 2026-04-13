package libgame

import (
	"fmt"
	"strings"
)

const _ActionCreate = "create"
const _ActionUpdate = "update"
const _ActionStart = "start"
const _ActionStartFork = "start-fork"

type Game interface {
	CreateServer()
	UpdateServer()
	StartServer()
	StartServerFork()
}

func RunAction(game Game, action string) (_badActionMessage string) {
	switch action {

	case _ActionCreate:
		game.CreateServer()

	case _ActionUpdate:
		game.UpdateServer()

	case _ActionStart:
		game.StartServer()

	case _ActionStartFork:
		game.StartServerFork()

	default:
		var sb strings.Builder
		fmt.Fprintf(&sb, "Invalid action:\n")
		fmt.Fprintf(&sb, "- %v\n", action)
		fmt.Fprintf(&sb, "Valid actions:\n")
		fmt.Fprintf(&sb, "- %v\n", _ActionCreate)
		fmt.Fprintf(&sb, "- %v\n", _ActionUpdate)
		fmt.Fprintf(&sb, "- %v\n", _ActionStart)
		fmt.Fprintf(&sb, "- %v\n", _ActionStartFork)
		return sb.String()

	}

	return ""
}
