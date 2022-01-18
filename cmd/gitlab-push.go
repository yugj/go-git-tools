package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	msg := ""
	for i, v := range os.Args {

		if i == 0 {
			continue
		}
		msg += v + " "
	}

	runCommand("git", "add", ".")
	runCommand("git", "commit", "-a", "-m", msg)
	runCommand("git", "push", "origin", "master")

}

// run and output
func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	log.Println("=>", string(output))
	if err != nil {
		log.Fatalln(err)
	}
}
