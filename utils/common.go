package utils

import (
	"fmt"
	"number-picker/utils/file"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetCurrPath 获取程序运行路径
func GetCurrPath() string {
	f, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(f)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

// MkDir ...
func MkDir(path string) string {
	path = filepath.FromSlash(path)
	// 如果是文件 并且文件存在 会获取文件路径 直接返回
	if file.IsFile(path) {
		path = file.GetFileDir(path)
		return path
	}

	//如果不是文件并且不存在 如果是含有带“.”的路径,一定要确保是路径而不是文件名,
	//所以确保传入的要么是有效的文件路径，要么是文件夹，不要传入不存在的文件路径
	if !file.IsExist(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic("创建目录失败: " + err.Error())
		}
	}
	return path
}

// WriteToFile ...
func WriteToFile(content string, filepath string, overwrite bool) error {
	if !overwrite {
		// 文件存在则返回错误
		if _, err := os.Stat(filepath); err == nil {
			return fmt.Errorf("该文件存在于此路径（你可以选择覆盖写入该文件）: %q", filepath)
		}
	}
	err := os.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	return nil
}
