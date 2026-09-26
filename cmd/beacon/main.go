package main

import (
	"beacon/internal/runner"
	"flag"
	"fmt"
	"os"
)

func main() {
	commentPrefix := flag.String("comment", "//", "The comment token used (default: //)")
	executablePath := flag.String("executable", "", "The path to the compiler executable (required)")
	testDirPath := flag.String("dir", "", "The path to the test directory (required)")

	flag.Parse()

	if *executablePath == "" {
		fmt.Println("Error: --executable is required.")
		flag.Usage()
		os.Exit(1)
	}

	if *testDirPath == "" {
		fmt.Println("Error: --dir is required.")
		flag.Usage()
		os.Exit(1)
	}

	cfg := runner.Config{
		ExecutablePath: *executablePath,
		CommentPrefix:  *commentPrefix,
		TestDirPath:    *testDirPath,
	}

	exitCode := runner.RunTests(cfg)
	os.Exit(exitCode)
}
