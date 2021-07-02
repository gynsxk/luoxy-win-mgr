package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"luoxy.xyz/winmgr/plugins"
)

const (
	BUF_LEN = 1024
)

var validPath = regexp.MustCompile("^/(execute)$")

func execute(w http.ResponseWriter, r *http.Request) {
	log.Println("receive new task")
	buf := make([]byte, r.ContentLength)
	r.Body.Read(buf)
	//fmt.Fprintf(w, "%s\n", buf)

	var tasks []plugins.Task
	err := json.Unmarshal(buf, &tasks)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err.Error())
		return
	}
	//fmt.Fprintf(w, "%+v", tasks)

	for _, task := range tasks {
		log.Printf("%+v", task)
		action := task["action"]

		fmt.Fprintf(w, "%+v\n", action)
		switch action {
		case "download":
			plugins.DownloadFileTask(task)

		case "cmd":

		case "runas":

		}
	}

}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
func worker() {
	http.HandleFunc("/execute", makeHandler(execute))
	http.ListenAndServe(":8080", nil)
}

// Run launches the service
func Run() error {
	log.Println("App RUN ...")

	// 接受命令，解释执行
	// 1. 下载可执行程序
	// 2. 执行脚本命令

	go worker()
	return nil
}
