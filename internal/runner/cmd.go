package runner

import (
	"beacon/internal/parser"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type ExecutionResult struct {
	Cmd    *exec.Cmd
	Output []byte
}

func runFileWithTimeout(cmd *exec.Cmd, timeout time.Duration) error {
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to execute beacon file: %w", err)
	}

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()

	select {
	case <-done:
		return nil
	case <-time.After(timeout):
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("failed to terminate beacon execution: %v", err)
		}
		return fmt.Errorf("beacon execution timed out: %s", cmd.Path)
	}
}

func executeFile(path string, args []string, settings parser.FileSettings) (*ExecutionResult, error) {
	cmd := exec.Command(path, args...)

	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &outBuffer

	// Run with your timeout logic
	err := runFileWithTimeout(cmd, settings.Timeout)
	if err != nil {
		return nil, err
	}

	result := &ExecutionResult{
		Cmd:    cmd,
		Output: outBuffer.Bytes(),
	}

	return result, nil
}
