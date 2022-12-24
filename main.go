package main

/*
*                                    Vidmoly Bot by Pix4
*                    Extract automatically the download links from vidmoly.me/dl
*
*                         Supports to extract data from multiple player links
*
										 Fully Proxied!
*							     Programmed by Z3NTL3 (aka Efdal)
*/

import (
	"Z3NTL3/Vidmoly-Bot/bot"
	"Z3NTL3/Vidmoly-Bot/builder"
	"Z3NTL3/Vidmoly-Bot/filesystem"
	"Z3NTL3/Vidmoly-Bot/globals"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type ProxyListGetter interface {
	getProxies(path string) ([]string, error) // any
}

type Sharpness struct {
	Proxies map[string] []string `yaml:"proxies"`
	filepath string
} 

type Context struct {
	proxyList []string // user:pass@ip:port proxies format
	vidmoly_cdn_link string // url
}

func (c Sharpness) getProxies() ([]string, error) {
	file, err := os.ReadFile(path.Join(c.filepath,"configuration","proxies.yaml"));if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &c); if err != nil {
		return nil,err
	}
	return c.Proxies["list"] ,nil
}

func errHandler(err error){
	fmt.Println("\033[31m",err, "\033[0m")
	os.Exit(-1)
}

func checkProxiesRegEx(ctx []string) (bool, []string){
	proxyRegex := regexp.MustCompile(`^(?P<username>[^:]+):(?P<password>[^@]+)@(?P<ip>[^:]+):(?P<port>\d+)$`)
	validsExists := false

	var valids []string

	for i, val := range ctx {
		valid := proxyRegex.MatchString(val)
		var value string

		if valid { 
			valids = append(valids, val)
			validsExists = true
			value = "\033[38;5;46mVALID\033[0m"
		} else {
			value = "\033[38;5;196mINVALID\033[0m"
		}
        builder.LLog(i,val,value)
	}
	return validsExists, valids
}

func validArgs(args *[]string) (valid bool){
	valid = true

	if(len(*args) == 2){
		fileREGEX := regexp.MustCompile(`^.*\.txt$`)
		numberREGEX := regexp.MustCompile(`^[0-9]+$`)

		fileArg := (*args)[0]
		timeoutArg := (*args)[1]

		if(!fileREGEX.MatchString(fileArg) || !numberREGEX.MatchString(timeoutArg)){
			valid = false
		}

	} else  {
		valid = false
	}
	return
}

func Init() ([]string, []string){
	cwd, err := os.Getwd(); if err != nil {
		errHandler(err)
	}
	basePath ,err := filepath.Abs(cwd);  if err != nil {
		errHandler(err)
	}

	builder.Logo()

	cliArgs := os.Args[1:]
	_valid := validArgs(&cliArgs); if(!_valid){
		builder.Log("INFO", "Invalid CLI arguments! See 'USAGE.md' file!", "Arguments", string(typedefs.Red),"")
		os.Exit(-1)
	}
	tOut, err := strconv.Atoi(cliArgs[2]); if (err != nil){
		builder.Log("INFO", "Invalid CLI arguments! See 'USAGE.md' file!", "Arguments", string(typedefs.Red),"")
		builder.Log("ERROR", err.Error(), "Arguments", string(typedefs.Red),"")
	}
	globals.Timeout = tOut

	var api Sharpness
	api.filepath = basePath

	proxies, err := api.getProxies(); if(err != nil){
		errHandler(err)
	}

	valids,proxies := checkProxiesRegEx(proxies); if(!valids){
		builder.Log("INFO", "Bad Proxy Format! Only username:pass@ip:port", "Proxy Format", string(typedefs.Red),"")
		os.Exit(-1)
	}

	webList := filesystem.Read(path.Join(basePath, cliArgs[0]))
	return strings.Split(webList, "\n"), proxies
}

func main() {
	bot.InitBypass(&[]string{"test"},&[]string{"test"},func() {
		fmt.Println()
	})
	// max_worker_count := runtime.NumCPU()
	// free_cores := 3
	
	// group := new(errgroup.Group)
	// group.SetLimit(10000 * (max_worker_count - free_cores))

	// webLinks, _ := Init()
	// websitesValidity := checker.Website(webLinks); if(!websitesValidity){
	// 	builder.Log("INFO", "Invalid web-links provided!", "Invalid URI", string(typedefs.Red),"")
	// }
}