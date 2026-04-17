package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	vaultDir    = "vault"
	hugoPostDir = "content/posts"
)

type FrontMatter struct {
	Published bool `yaml:"publish"`
}

func parseFrontMatter(path string) (FrontMatter, error) {
	var fm FrontMatter

	content, err := os.ReadFile(path)
	if err != nil {
		return fm, err
	}

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return fm, nil
	}

	yaml.Unmarshal([]byte(parts[1]), &fm)
	return fm, nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func main() {
	os.MkdirAll(hugoPostDir, 0755)

	entries, err := os.ReadDir(vaultDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading vault dir: %v\n", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		src := filepath.Join(vaultDir, entry.Name())
		fm, err := parseFrontMatter(src)
		if err != nil {
			fmt.Printf("Skipped (parse error): %s\n", entry.Name())
			continue
		}

		if fm.Published {
			dst := filepath.Join(hugoPostDir, entry.Name())
			if err := copyFile(src, dst); err != nil {
				fmt.Printf("Error copying %s: %v\n", entry.Name(), err)
			} else {
				fmt.Printf("Copied: %s\n", entry.Name())
			}
		} else {
			fmt.Printf("Skipped: %s\n", entry.Name())
		}
	}
}
