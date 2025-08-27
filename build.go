//go:build exclude

package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

type target struct {
    os      string
    arch    string
    ext     string
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
    now := time.Now().UTC().Format("02Jan06")

    fmt.Println("removing old binaries")
    _ = os.RemoveAll("bin")
    _ = os.Mkdir("bin", 0o755)

    for _, t := range targets {
        outName := fmt.Sprintf("CTCloniaTextures-%s-%s-%s%s",
            t.os, t.arch, now, t.ext)
        outPath := filepath.Join("bin", outName)

        fmt.Printf("building %s %s\n", t.os, t.arch)
        cmd := exec.Command("go", "build", "-trimpath", "-o", outPath, ".")
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
