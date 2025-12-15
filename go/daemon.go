package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	// A special environment variable to signal the child process.
	daemonEnvVar = "GO_DAEMON_CHILD"
	logFile      = "daemon.log"
)

// This is the actual long-running task for the daemon.
func worker() {
	// Open a log file for the daemon to write to.
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Log a message every 10 seconds to show it's alive.
	for {
		log.Println("Daemon is working...")
		time.Sleep(10 * time.Second)
	}
}

func main() {
	// Check if the special environment variable is set.
	// If it is, this is the child process, so we start the worker.
	if os.Getenv(daemonEnvVar) == "1" {
		worker()
		return
	}

	// --- This is the parent process ---

	// Get the path of the currently running executable.
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf("Cannot find executable: %v", err)
	}

	// Prepare the command to re-run the executable.
	cmd := exec.Command(executable)

	// Set the special environment variable for the child process.
	// This is the signal for the child to become the daemon.
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=1", daemonEnvVar))

	// Detach the process from the current terminal.
	// This is the most critical part for daemonization on Unix-like systems.
	// Setting Setsid to true creates a new session, detaching the child.
	cmd.SysProcAttr = &os.SysProcAttr{
		Setsid: true,
	}

	// Start the command. We don't wait for it to complete.
	err = cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start daemon: %v", err)
	}

	// Parent process prints the child's PID and exits.
	fmt.Printf("Daemon started with PID: %d\n", cmd.Process.Pid)
	fmt.Printf("Logs will be written to %s\n", logFile)
	os.Exit(0)
}
