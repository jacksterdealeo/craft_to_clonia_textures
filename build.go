//go:build exclude

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type target struct {
	os   string
	arch string
	ext  string
}

var targets = []target{
	{"windows", "386", ".exe"},
	{"windows", "amd64", ".exe"},
	{"darwin", "amd64", ""},
	{"darwin", "arm64", ""},
	{"linux", "386", ""},
	{"linux", "amd64", ""},
	{"linux", "arm64", ""},
}

func main() {
	now := []byte(time.Now().UTC().Format("02jan06"))
	tag, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		tag = []byte(string(now) + "noTag")
	} else {
		tag = []byte(strings.TrimRight(string(tag), "\n"))
	}
	tagString := string(tag)

	fmt.Println("removing old binaries")
	_ = os.RemoveAll("bin")
	_ = os.Mkdir("bin", 0o755)

	for _, t := range targets {
		outName := fmt.Sprintf("CTCloniaTextures-%s-%s-%s%s",
			t.os, t.arch, tagString, t.ext)
		outPath := filepath.Join("bin", outName)

		fmt.Printf("building %s %s\n", t.os, t.arch)
		ldflags := fmt.Sprintf("-X main.version=%s", strings.TrimSpace(tagString))

		cmd := exec.Command("go", "build",
			"-trimpath",
			"-ldflags", ldflags,
			"-o", outPath, ".")

		cmd.Env = append(os.Environ(),
			"GOOS="+t.os,
			"GOARCH="+t.arch,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "build failed for %s/%s: %v\n", t.os, t.arch)
			os.Exit(1)
		}
	}

	fmt.Println("build tasks completed")
}
