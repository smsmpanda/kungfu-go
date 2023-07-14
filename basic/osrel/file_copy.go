package osrel

import (
	"fmt"
	"io"
	"os"
)

func Gorun123rg() {
	CopyFile("resources/target.txt", "resources/source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// 注意 defer 的使用：当打开 dst 文件时发生了错误，那么 defer 仍然能够确保 src.Close() 执行。如果不这么做，src 文件会一直保持打开状态并占用资源。
