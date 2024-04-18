package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

// args = {program name} {command} {args}

func main() {
	args := os.Args

	if len(args) < 2 {
		showIntro()
		return
	}

	command := args[1]
	switch command {
	case "push":
		handlePush(args)
	case "js":
		handleJs()
	default:
		color.Red("Unknown command: %s", command)
		os.Exit(1)
	}
}

func showIntro() {
	color.Green("Usage: cast [command] [args]")
	fmt.Println()
	color.Green("Commands:")
	fmt.Println()
	fmt.Println("- push [message] - add files, commit and pushes to git")
}

func runSystemCommand(name string, args ...string) {
	msg := "Executing: " + name
	for _, arg := range args {
		msg = msg + " " + arg
	}
	color.Yellow(msg)

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}
}

func handlePush(args []string) {
	if len(args) < 3 {
		color.Red("Error: Message is required as third argument")
		fmt.Println("Example: cast push \"update readme\"")
		os.Exit(1)
	}

	message := "\"" + args[2] + "\""
	runSystemCommand("git", "add", ".")
	runSystemCommand("git", "commit", "-m", message)
	runSystemCommand("git", "push")
}

func handleJs() {
	if _, err := os.Stat("pnpm-lock.yaml"); os.IsNotExist(err) {
		runSystemCommand("npm", "run", "dev")
	} else {
		runSystemCommand("pnpm", "run", "dev")
	}
}
