package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var extensionMap = map[string]string{
	"mp4":   "video",
	"jpg":   "image",
	"png":   "image",
	"jpeg":  "image",
	"gif":   "image",
	"mp3":   "audio",
	"flv":   "video",
	"mov":   "video",
	"avi":   "video",
	"mkv":   "video",
	"wmv":   "video",
	"html":  "web",
	"css":   "web",
	"js":    "web",
	"cpp":   "code",
	"py":    "code",
	"go":    "code",
	"rs":    "code",
	"ipynb": "notebook",
	"pdf":   "document",
	"doc":   "document",
	"docx":  "document",
	"xls":   "document",
	"xlsx":  "document",
	"ppt":   "document",
	"pptx":  "document",
	"txt":   "document",
	"zip":   "archive",
	"rar":   "archive",
	"7z":    "archive",
	"iso":   "iso",
	"":      "other",
}

func getFileType(filename string) string {
	idx := strings.LastIndex(filename, ".")
	var result string
	if idx == -1 {
		result = ""
	} else {
		result = filename[idx+1:]
	}
	return result
}

func moveFileToFolder(root string, filename string, folder string) {
	if _, err := os.Stat(root + folder); os.IsNotExist(err) {
		os.Mkdir(root+folder, 0755)
	}
	// check if the file with the same nema is already in the folder

	if _, err := os.Stat(root + folder + "/" + filename); err == nil {
		// error panick
		fmt.Println("File " + filename + " already exists in " + folder + "/")
		return
	}
	err := os.Rename(root+filename, root+folder+"/"+filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Moved " + filename + " to " + folder + "/")
}

func listFiles(path string) {
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {

	path := "/home/anasr/Downloads/"

	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileType := getFileType(file.Name())
		folder := extensionMap[fileType]
		moveFileToFolder(path, file.Name(), folder)
	}

}
