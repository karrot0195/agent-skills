package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	target := flag.String("Target", "", "target directory (required)")
	flag.Parse()

	if *target == "" {
		fmt.Fprintln(os.Stderr, "Error: -Target is required")
		flag.Usage()
		os.Exit(1)
	}

	src := "skills"
	if _, err := os.Stat(src); err != nil {
		fmt.Fprintf(os.Stderr, "Error: source %s not found: %v\n", src, err)
		os.Exit(1)
	}

	if err := copyDir(src, *target); err != nil {
		fmt.Fprintf(os.Stderr, "Error copying skills: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Copied %s to %s\n", src, *target)
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}

		return dstFile.Chmod(info.Mode())
	})
}
