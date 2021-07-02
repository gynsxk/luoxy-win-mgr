package common

import (
	"fmt"
	"log"
	"os"
	"time"
)

type RotatorWriter struct {
	prefix         string
	filename       string
	historySize    int
	file           *os.File
	fileCreateTime time.Time
}

func NewRotatorWriter(filename string) *RotatorWriter {
	w := &RotatorWriter{}
	w.prefix = filename
	w.filename = fmt.Sprintf("%s%s%c%s.log", GetWorkDir(), "logs", os.PathSeparator, w.prefix)
	w.historySize = 7
	w.file = nil
	return w
}

func (w *RotatorWriter) Write(p []byte) (n int, err error) {
	if w.file == nil {
		// 获取最后修改时间
		stat, err := os.Stat(w.filename)
		if err != nil {
			w.fileCreateTime = time.Now()
		} else {
			// 没有创建时间，这个可以凑活使用
			w.fileCreateTime = stat.ModTime()
		}

		file, err := os.OpenFile(w.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		w.file = file
	}

	day := time.Now().Day()
	if day != w.fileCreateTime.Day() {
		w.file.Close()
		dateStr := w.fileCreateTime.Format("2006-01-02")
		dstFileName := fmt.Sprintf("%s%s%c%s-%s%s", GetWorkDir(), "logs", os.PathSeparator, w.prefix, dateStr, ".log")
		os.Rename(w.filename, dstFileName)

		file, err := os.OpenFile(w.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		w.fileCreateTime = time.Now()
		w.file = file
	}
	return w.file.Write(p)
}
