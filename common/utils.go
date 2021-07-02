package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var workDir = ""

func GetWorkDir() string {
	if len(workDir) > 0 {
		return workDir
	}
	programPath, err2 := os.Executable()
	if err2 != nil {
		log.Fatal("获取可执行程序路径失败")
	}
	workDir = fmt.Sprintf("%s%c", filepath.Dir(programPath), os.PathSeparator)
	return workDir
}
