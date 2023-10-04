package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// go run cmd/main.go

func main() {
	// 定时删除的间隔时间
	interval := 1 * time.Second

	// 获取当前可执行文件的路径
	exePath := GetRunPath()

	fmt.Printf("Executable path: %s\n", exePath)
	// 获取要删除文件的文件夹路径
	folderPath := filepath.Join(exePath, "test")
	fmt.Printf("folderPath path: %s\n", folderPath)

	// 启动定时任务
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 执行定时删除操作
			deleteFiles(folderPath)
		}
	}
}

// GetRunPath 获取执行目录作为默认目录
func GetRunPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currentPath
}
func deleteFiles(folderPath string) {
	// 打开要删除的文件夹
	dir, err := os.Open(folderPath)
	if err != nil {
		fmt.Println("Error opening folder:", err)
		return
	}
	defer dir.Close()

	// 读取文件列表
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading folder contents:", err)
		return
	}

	// 当前时间
	now := time.Now()

	// 删除文件
	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())
		fileModTime := file.ModTime()

		// 比较文件的修改时间，如果早于当前时间则删除
		if fileModTime.Before(now) {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("Error deleting file %s: %v\n", filePath, err)
			} else {
				fmt.Printf("Deleted file: %s\n", filePath)
			}
		}
	}
}
