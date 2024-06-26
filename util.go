package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/s00500/env_logger"
)

func main() {
	// CheckArgs("<directory>")
	// directory := os.Args[1]
	// // Opens an already existing repository.
	// r, err := git.PlainOpen(directory)
	// log.Should(err)

	// w, err := r.Worktree()
	// log.Should(err)

	// status, err := w.Status()
	// log.Should(err)

	// fmt.Println(status)
	log.Info("start")
	cmd := exec.Command("git", "describe --tags")

	// Get the output
	output, err := cmd.Output()
	if err != nil {
		log.Errorf("Error executing command:", err)
		return
	}

	// Print the output
	fmt.Println(string(output))
}

func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		log.Warnf("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}
