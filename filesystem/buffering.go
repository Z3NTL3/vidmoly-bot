package filesystem

import (
	"Z3NTL3/Vidmoly-Bot/builder"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"bufio"
	"os"
)


func Read(filepath string) (string){
	file, err := os.Open(filepath); if(err != nil){
		builder.Log("INFO", "Cannot open file, does it exist?", "FileSystem",string(typedefs.Red))
		os.Exit(-1)
	}
	
	defer file.Close()
	buff := bufio.NewReader(file)
	data := ""

	for {
		chunk := make([]byte, 1042 * 4)
		rLen,err := buff.Read(chunk); if(err != nil ){
			break
		}
		data += string(chunk[:rLen])
	}
	return data
}