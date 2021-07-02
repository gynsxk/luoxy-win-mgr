package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	baseFolder = "dist/"
	dst     = `dist/winmgr.zip`
	zipBase = "winmgr/"
)

func main() {
	ZipWriter()
}
func ZipWriter() {
	

	// Get a Buffer to Write To
	outFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, baseFolder, zipBase)

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}
func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".zip") {
				continue
			}
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {
			// Recurse
			header := &zip.FileHeader{
				Name:   baseInZip + file.Name() + "/",
				Method: zip.Deflate,
			}
			header.SetMode(os.ModeDir)
			headerX, err := w.CreateHeader(header)
			if err != nil {
				fmt.Println(err)
			}
			b := make([]byte, 0)
			headerX.Write(b)

			newBase := basePath + file.Name() + "/"

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
