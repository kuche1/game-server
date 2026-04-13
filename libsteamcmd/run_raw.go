package libsteamcmd

import (
	"log"
	"os"
	"os/exec"
)

func runCustom(forceInstallDir string, additionalCmdline ...string) {
	// TODO: we MAY need to use `/usr/bin/steamcmd` instead, need to test this
	executable := "steamcmd"

	///// check that steamcmd is installed

	cmd := exec.Command(executable, "--help")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	///// run command

	args := make([]string, 0)
	args = append(args, "+force_install_dir", forceInstallDir)
	args = append(args, "+login", "anonymous")
	args = append(args, additionalCmdline...)
	args = append(args, "+exit")

	cmd = exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
