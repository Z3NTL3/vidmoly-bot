package filesystem

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func WriteToSaveFile(data string) (error){
	if len(data) != 1 || data != ""{
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
		_, err = file.WriteString(fmt.Sprintf("%s\n",data)); if err != nil {
			return err
		}
	}

	return nil
}