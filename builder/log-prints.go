package builder

import "fmt"

func Log(log_type string, info string, errType string, color string, delim string) {

	fmt.Printf("\n\033[1m\033[38;5;134m%s: \033[0m\033[1m[\033[0m%s\033[1m%s\033[0m\033[1m]\033[0m \033[1m-->\033[0m \033[1m%s%s\033[0m%s", log_type,color, errType, color,info,delim)
}

func LLog(i int,val string, value string){
	fmt.Printf("\033[1m\033[38;5;134mChecking proxy format: \033[0m\033[1m[\033[0m\033[38;5;216mID PROXY: %d\033[0m\033[1m]\033[0m \033[1m-->\033[0m \033[38;5;199m\033[1m%s\033[0m - \033[1m%v \033[0m\r\n", (i+1) , val, value)
}