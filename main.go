package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func analyzeDirectory(dirPath string) error {
    return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            return nil
        }

        // 파일 메타데이터 분석
        fmt.Printf("File: %s\n", path)
        fmt.Printf("Size: %d bytes\n", info.Size())
        fmt.Printf("Last Modified: %s\n", info.ModTime())

        // 삭제된 파일 흔적 찾기 (예: 파일명에 '._' 접두사가 있는 경우)
        if filepath.Base(path)[:2] == "._" {
            fmt.Println("This might be a remnant of a deleted file.")

            // 복구 가능성 평가 (예시로, 파일 크기가 0이 아니면 복구 가능하다고 가정)
            if info.Size() > 0 {
                fmt.Println("Recovery might be possible.")
            } else {
                fmt.Println("Recovery is unlikely.")
            }
        } else {
            fmt.Println("This is not a deleted file.")
        }

        fmt.Println(strings.Repeat("-", 50))
        return nil
    })
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run script.go <directory_path>")
        os.Exit(1)
    }

    dirPath := os.Args[1]
    err := analyzeDirectory(dirPath)
    if err != nil {
        fmt.Printf("Error analyzing directory: %v\n", err)
    }
}
