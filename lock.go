// This code is use to prevent Windows from locking out with the use of PowerShell commands to simulate key presses.


package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func simulateKeyPress(keyCode string) {
	psScript := fmt.Sprintf(`[System.Windows.Forms.SendKeys]::SendWait("{%s}")`, keyCode)
	cmd := exec.Command("powershell", "-Command", psScript)
	cmd.Run()
}

func noLock(keyCode string) {
	fmt.Println("Press CTRL+C to stop.")

	for {
		simulateKeyPress(keyCode)

		time.Sleep(2 * time.Second)
	}
}

func main() {
	fmt.Println("\nPrevent Windows screen lock")
	fmt.Print("Enter key code to simulate (e.g., C): ")

	var keyCode string
	_, err := fmt.Scan(&keyCode)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	// Capture Ctrl+C to stop the program gracefully
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	go noLock(keyCode)

	fmt.Println("Running")

	// Wait for an interrupt signal (Ctrl+C)
	<-interrupt

	fmt.Println("\nStopped")
}
