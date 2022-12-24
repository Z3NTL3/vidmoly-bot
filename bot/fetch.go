package bot

import (
	"Z3NTL3/Vidmoly-Bot/config"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"
)
func InitBypass(webList, proxies *[]string, callb typedefs.BypassType){
	transport := config.Config()
	client := http.Client{
		Transport: transport,
	}
	var bodyReader io.ReadCloser
	req, err := http.NewRequestWithContext(context.Background(), "GET", "https://vidmoly.me/d/80sqn2j71v26", bodyReader); if(err != nil){
		fmt.Println(err)
	}

	resp, err := client.Do(req); if(err != nil){
		fmt.Println(err)
	}

	var length int
	var Storage bytes.Buffer
	
	for {
		buffer := make([]byte,1042)
		readLenght, err := io.ReadFull(resp.Body,buffer); if(err != nil){
			if(err == io.EOF){
				break
			}
		}
		length += readLenght
		Storage.Write(buffer)
	}
	
	body := make([]byte,length)
	length, _ = Storage.Read(body)
	fmt.Println(string(body[0:length]))

	fmt.Println(config.TLS_Vers[resp.TLS.Version])
	// fmt.Println(body)


	// doc,err := htmlquery.Parse(strings.NewReader(body))
	// nodes := htmlquery.Find(doc, "/html/body/div/div[2]/p/code[1]")
	// if err != nil {
	// 	panic(`not a valid XPath expression.`)
	// }
	// fmt.Println(htmlquery.InnerText(nodes[0]))
}