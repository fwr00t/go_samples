package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func simulateKeyPress(keyCode string) {
	psScript := fmt.Sprintf(`Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.SendKeys]::SendWait("%s")`, keyCode)
	cmd := exec.Command("powershell", "-Command", psScript)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing PowerShell command:", err)
	}
}

func noLock(keyCode string) {
	fmt.Println("Press CTRL+C to stop.")

	for {
		simulateKeyPress(keyCode)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	fmt.Println("\nPrevent Windows screen lock")
	fmt.Print("Enter key code to simulate (e.g., C): ")

	reader := bufio.NewReader(os.Stdin)
	keyCode, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	// Capture Ctrl+C to stop the program gracefully
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	go noLock(keyCode)

	fmt.Println("Running. . . . .")

	// Wait for an interrupt signal (Ctrl+C)
	<-interrupt

	fmt.Println("\nStopped")
}
