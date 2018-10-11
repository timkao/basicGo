package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copyFile("target.txt", "source.txt")
	fmt.Println("file is copied")
}

func copyFile(dst string, src string) (written int64, err error) {
	srcFile, err1 := os.Open(src)
	if err1 != nil {
		fmt.Println("fail to open", src)
		return
	}
	defer srcFile.Close()

	dstFile, err2 := os.Create(dst) // need to use "create". if the file already exists, we should write to it instead of copy
	if err2 != nil {
		fmt.Println("fail to open", dst)
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)

}
