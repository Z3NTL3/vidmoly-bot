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
	"Z3NTL3/Vidmoly-Bot/bot/fetch"
	"Z3NTL3/Vidmoly-Bot/builder"
	"Z3NTL3/Vidmoly-Bot/checker"
	"Z3NTL3/Vidmoly-Bot/filesystem"
	"Z3NTL3/Vidmoly-Bot/globals"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

type ProxyListGetter interface {
	getProxies(path string) ([]string, error) // any
}

type Sharpness struct {
	Proxies map[string] []string `yaml:"proxies"`
	filepath string
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

func Init() ([]string, []string, error){
	cwd, err := os.Getwd(); if err != nil {
		globals.ErrHandler(err)
	}
	basePath ,err := filepath.Abs(cwd);  if err != nil {
		globals.ErrHandler(err)
	}

	builder.Logo()

	cliArgs := os.Args[1:]
	_valid := validArgs(&cliArgs); if(!_valid){
		builder.Log("INFO", "Invalid CLI arguments! See 'USAGE.md' file!", "Arguments", string(typedefs.Red),"")
		os.Exit(-1)
	}
	if(len(cliArgs) == 3){
		tOut, err := strconv.Atoi(cliArgs[2]); if (err != nil){
			builder.Log("INFO", "Invalid CLI arguments! See 'USAGE.md' file!", "Arguments", string(typedefs.Red),"")
			builder.Log("ERROR", err.Error(), "Arguments", string(typedefs.Red),"")
		}		
		globals.Timeout = tOut
	}

	var api Sharpness
	api.filepath = basePath

	proxies, err := api.getProxies(); if(err != nil){
		globals.ErrHandler(err)
	}

	valids,proxies := checkProxiesRegEx(proxies); if(!valids){
		builder.Log("INFO", "Bad Proxy Format! Only username:pass@ip:port or you have 0 proxies", "Proxy Format", string(typedefs.Red),"")
		os.Exit(-1)
	}

	webList := filesystem.Read(path.Join(basePath, cliArgs[0]))
	if(len(webList) == 0 || len(proxies) == 0){
		return []string{"empty"}, []string{"empty"}, nil
	}
	return strings.Split(webList, "\n"), proxies, nil
}

func main() {
	infos := make(map[string] []string)

	max_worker_count := runtime.NumCPU()
	free_cores := 3
	
	group := new(errgroup.Group)
	group.SetLimit(10000 * (max_worker_count - free_cores))

	webLinks, proxies, err:= Init(); if(err != nil){
		builder.Log("INFO", err.Error(), "Error", string(typedefs.Red),"")
		os.Exit(-1)
	}

	infos["links"] = webLinks
	infos["proxies"] = proxies
	
	websitesValidity := checker.Website(webLinks); if(!websitesValidity || webLinks[0] == "empty"){
		builder.Log("INFO", "Invalid web-links provided or there is no links data in your file present", "Invalid URI", string(typedefs.Red),"")
	}
	// fetch.InitBypass(&a,&b,func() {
	// 	fmt.Println() // dit moet allemaal eigenlijk gestart worden met waitgroup of errgroup voor goroutines dat doe ik later dit is nog in test fase
	// })

	valid := checker.CheckInfo(infos); if(!valid){
		builder.Log("INFO", "There is no data present in proxies.yaml or your provided weblist txt file", "No Data", string(typedefs.Red),"")
	}

	for i,v := range infos["links"] {
		group.Go( func()(error){
			return  fetch.InitBypass(i+1,&v,&infos["proxies"][0], bot.Bypass)
		})
	}

	err = group.Wait(); if(err != nil){
		builder.Log("INFO", err.Error(), "Error", string(typedefs.Red),"")
	}
}