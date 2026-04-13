package libutil

import (
	"log"
	"os"
	"os/exec"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatal(err)
	}

	return true
}

func IsFolder(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatal(err)
	}

	return info.IsDir()
}

func Exec(executable string, printOutput bool, cwd string, blocking bool, args ...string) {
	cmd := exec.Command(executable, args...)

	cmd.Dir = cwd

	if printOutput {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	var err error

	if blocking {
		err = cmd.Run()
	} else {
		err = cmd.Start()
	}

	if err != nil {
		log.Fatal(err)
	}
}
