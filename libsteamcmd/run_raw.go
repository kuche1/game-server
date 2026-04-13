package libsteamcmd

import (
	"github.com/kuche1/game-server/libutil"
)

func runCustom(forceInstallDir string, additionalCmdline ...string) {
	// TODO: we MAY need to use `/usr/bin/steamcmd` instead, need to test this
	executable := "steamcmd"

	///// check that steamcmd is installed

	libutil.Exec(
		executable,
		false,
		"",
		true,
		"--help",
	)

	///// run command

	args := make([]string, 0)
	args = append(args, "+force_install_dir", forceInstallDir)
	args = append(args, "+login", "anonymous")
	args = append(args, additionalCmdline...)
	args = append(args, "+exit")

	libutil.Exec(
		executable,
		true,
		"",
		true,
		args...,
	)
}
