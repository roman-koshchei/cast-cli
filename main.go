package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cast <command> [args]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "push":
		if len(os.Args) < 3 {
			fmt.Println("Usage: cast push <message>")
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
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error adding files:", err)
		fmt.Println("Output:", string(output))
		return err
	}
	fmt.Println("Added files successfully")
	fmt.Println("Output:", string(output))

	cmd = exec.Command("git", "commit", "-m", message)
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		fmt.Println("Output:", string(output))
		return err
	}
	fmt.Println("Committed changes successfully")
	fmt.Println("Output:", string(output))

	cmd = exec.Command("git", "push")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error pushing changes:", err)
		fmt.Println("Output:", string(output))
		return err
	}
	fmt.Println("Pushed changes successfully")
	fmt.Println("Output:", string(output))

	fmt.Println("Pushed changes with message:", message)
	return nil
}
