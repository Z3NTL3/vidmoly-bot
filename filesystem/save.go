package filesystem

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func WriteToSaveFile(data string) (error){
	data = strings.ReplaceAll(data, "\n","")
	data = strings.ReplaceAll(data, "\r","")

	if len(data) != 0 && data != "" || strings.Contains(data,"mp4"){
		cwd, err := os.Getwd(); if err != nil {
			return err
		}
		basePath ,err := filepath.Abs(cwd);  if err != nil {
			return err
		}
		file,err := os.OpenFile(path.Join(basePath,"save","goods.txt"),os.O_CREATE|os.O_APPEND, 0644); if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("%s\r\n",data)); if err != nil {
			return err
		}
	}

	return nil
}