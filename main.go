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
	}

	// command := os.Args[1]
	// switch command {
	// case "push":
	// 	handlePush(os.Args)
	// case "js":
	// 	err := runJSDev()
	// 	if err != nil {
	// 		color.Red("Error running dev server: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// default:
	// 	color.Red("Unknown command: %s\n", os.Args[1])
	// 	os.Exit(1)
	// }
}

func showIntro() {
	color.Green("Usage: cast [command] [args]")
	fmt.Println()
	color.Green("Commands:")
	fmt.Println()
	fmt.Println("- push [message] - add files, commit and pushes to git")
}

func runSystemCommand(name string, args ...string) bool {
	fmt.Printf("Executing: %s ", name)
	for _, arg := range args {
		fmt.Print(arg)
	}
	fmt.Println()

	output, err := exec.Command(name, args...).CombinedOutput()
	if len(output) > 0 {
		fmt.Println("Output:")
		fmt.Println(output)
	}

	if err != nil {
		color.Red("Error: %v\n", err)
		return false
	}

	return true
}

func handlePush(args []string) {
	if len(args) < 3 {
		color.Red("Error: Message is required as third argument")
		fmt.Println("Example: cast push \"update readme\"")
		os.Exit(1)
	}

	message := args[2]

	if !runSystemCommand("git", "add", ".") {
		return
	}
	if !runSystemCommand("git", "commit", "-m", message) {
		return
	}
	if !runSystemCommand("git", "push") {
		return
	}
}

func runJSDev() error {
	if _, err := os.Stat("pnpm-lock.yaml"); os.IsNotExist(err) {
		return runNPMDev()
	}
	return runPNPMDev()
}

func runNPMDev() error {
	cmd := exec.Command("npm", "run", "dev")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runPNPMDev() error {
	cmd := exec.Command("pnpm", "run", "dev")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
