package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func installLibrary(libraryName string) {
	// Check if Homebrew is installed
	_, err := exec.LookPath("brew")
	if err != nil {
		fmt.Println("Error: Homebrew is not installed. Please install Homebrew first.")
		os.Exit(1)
	}

	// Check if the library is already installed
	cmd := exec.Command("brew", "list", "--formula", libraryName)
	err = cmd.Run()
	if err == nil {
		fmt.Printf("%s is already installed.\n", libraryName)
	} else {
		// Install the library using Homebrew
		cmd = exec.Command("brew", "install", libraryName)
		err := cmd.Run()
		if err == nil {
			fmt.Printf("%s has been successfully installed.\n", libraryName)
		} else {
			fmt.Printf("Error: Failed to install %s.\n", libraryName)
			os.Exit(1)
		}
	}
}

func InstallDependencies(dependencyFile string) {
	file, err := os.Open(dependencyFile)
	if err != nil {
		fmt.Printf("Error: Unable to open %s\n", dependencyFile)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read each line from the file and install the library
	for scanner.Scan() {
		libraryName := scanner.Text()
		installLibrary(libraryName)
	}

}
