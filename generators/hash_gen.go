package generators

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// GetFileHash - Creates Hash where key is file name & value is file contents
func GetFileHash(folder1 string, extention string) map[string]string {
	var files []string
	Folder1Hash := make(map[string]string)
	err := filepath.Walk(folder1, func(curpath string, info os.FileInfo, err error) error {

		if info.IsDir() == false {
			curExtention := path.Ext(curpath)
			if curExtention == extention || extention == "*" {
				files = append(files, curpath)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePathArr := strings.Split(file, "\\")
		fileName := filePathArr[len(filePathArr)-1]
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}
		text := string(content)
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, " ", "", -1)
		Folder1Hash[fileName] = text
	}
	return Folder1Hash
}
