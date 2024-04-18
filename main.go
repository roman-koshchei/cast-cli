package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: cast <command> [args]")
	// 	os.Exit(1)
	// }

	for _, item := range os.Args {
		fmt.Println(item)
	}

	command := os.Args[1]
	switch command {
	case "push":
		handlePush(os.Args)
	case "js":
		err := runJSDev()
		if err != nil {
			color.Red("Error running dev server: %v\n", err)
			os.Exit(1)
		}
	default:
		color.Red("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func handlePush(args []string) {
	if len(args) < 3 {
		color.Red("Error: Message is required as third argument")
		fmt.Println("Example: cast push \"update readme\"")
		os.Exit(1)
	}

	message := strings.Join(os.Args[2:], " ")
	err := gitPush(message)
	if err != nil {
		color.Red("Error: %v\n", err)
		os.Exit(1)
	}
}

func gitPush(message string) error {
	cmd := exec.Command("git", "add", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		color.Red("Error adding files: %v\n", err)
		color.Red("Output: %s\n", string(output))
		return err
	}
	color.Green("Added files successfully\n")
	color.Green("Output: %s\n", string(output))

	cmd = exec.Command("git", "commit", "-m", message)
	output, err = cmd.CombinedOutput()
	if err != nil {
		color.Red("Error committing changes: %v\n", err)
		color.Red("Output: %s\n", string(output))
		return err
	}
	color.Green("Committed changes successfully\n")
	color.Green("Output: %s\n", string(output))

	cmd = exec.Command("git", "push")
	output, err = cmd.CombinedOutput()
	if err != nil {
		color.Red("Error pushing changes: %v\n", err)
		color.Red("Output: %s\n", string(output))
		return err
	}
	color.Green("Pushed changes successfully\n")
	color.Green("Output: %s\n", string(output))

	fmt.Println("Pushed changes with message:", message)
	return nil
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
