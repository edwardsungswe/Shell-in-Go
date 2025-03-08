package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Get and display the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting current directory:", err)
			continue
		}

		fmt.Print(cwd + "> ") // Show the correct working directory in the prompt

		// Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			// Handle EOF and error situations
			if err.Error() == "EOF" {
				// Exit the loop gracefully on EOF
				fmt.Println("\nExiting shell...")
				break
			} else {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
		}

		// Trim any extra whitespace
		input = strings.TrimSpace(input)

		// If the input is empty, continue to the next iteration
		if input == "" {
			continue
		}

		// Handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Split the input into command and arguments
	args := strings.Split(input, " ")

	// Check for built-in commands
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		// Change the directory
		err := os.Chdir(args[1])
		if err != nil {
			return fmt.Errorf("failed to change directory: %v", err)
		}
		return nil
	case "exit":
		os.Exit(0)
	}

	// Pass the program and arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output devices
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return any error
	return cmd.Run()
}
