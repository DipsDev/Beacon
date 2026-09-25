package main

import (
	"beacon/internal"
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
		fmt.Println("Error: --executable-path (-x) is required.")
		flag.Usage()
		os.Exit(1)
	}

	if *testDirPath == "" {
		fmt.Println("Error: --test-dir (-t) is required.")
		flag.Usage()
		os.Exit(1)
	}

	cfg := internal.Config{
		ExecutablePath: *executablePath,
		CommentPrefix:  *commentPrefix,
		TestDirPath:    *testDirPath,
	}

	exitCode := internal.RunTests(cfg)
	os.Exit(exitCode)
}
