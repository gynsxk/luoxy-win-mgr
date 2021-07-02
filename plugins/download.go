package plugins

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"luoxy.xyz/winmgr/common"
)

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadFileTask(task Task) error {
	url := fmt.Sprintf("%v", task["url"])
	fullPath := filepath.Join(common.GetWorkDir(), "app", fmt.Sprintf("%v", task["dir"]))
	os.MkdirAll(fullPath, 0777)
	fullPath = fullPath + string(os.PathSeparator) + fmt.Sprintf("%v", task["filename"])
	return downloadFile(fullPath, url)

}
