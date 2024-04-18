package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: do <command> [args]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "push":
		if len(os.Args) < 3 {
			fmt.Println("Usage: do push <message>")
			os.Exit(1)
		}
		message := strings.Join(os.Args[2:], " ")
		err := gitPush(message)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func gitPush(message string) error {
	cmd := exec.Command("git", "add", ".")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", message)
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "push")
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Pushed changes with message:", message)
	return nil
}