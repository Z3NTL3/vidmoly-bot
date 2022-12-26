package fetch

import (
	"Z3NTL3/Vidmoly-Bot/builder"
	"Z3NTL3/Vidmoly-Bot/checker"
	"Z3NTL3/Vidmoly-Bot/config"
	"Z3NTL3/Vidmoly-Bot/globals"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"Z3NTL3/Vidmoly-Bot/xpath"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/errors"
)
func InitBypass(task int,web, proxy *string, callb typedefs.BypassType) (error){
	rand.Seed(time.Now().Unix())
	err := checker.CheckProxy(proxy); if(err != nil){
		return err
	}

	transport,err := config.Config(*proxy); if err != nil {
		return err
	}
	client := http.Client{
		Transport: transport,
		Timeout: (time.Duration(globals.Timeout)*time.Second),
	}
	var bodyReader io.ReadCloser // auto close body
	req, err := http.NewRequest("GET", *web, bodyReader); if(err != nil){
		builder.Log("ERR",fmt.Sprintf("Could not init Req Obj for: %s",*web), "GET FETCH", string(typedefs.Red),"")
		builder.Log("Err Info",err.Error(), "Req obj init", string(typedefs.Red),"\n")
		return nil
	}
	req.Header.Add("User-Agent", typedefs.Headers_[rand.Intn(len(typedefs.Headers_))])
	
	resp, err := client.Do(req); if(err != nil){
		builder.Log("ERR",fmt.Sprintf("Could not fetch: %s",*web), "GET FETCH", string(typedefs.Red),"")
		builder.Log("Err Info",err.Error(), "GET FETCH", string(typedefs.Red),"\n")
		return nil
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
	bodyFull := string(body[0:length])
	ver := config.TLS_Vers[resp.TLS.Version]

	HTML_DOM := xpath.Document{Htmldoc: bodyFull}
	origin, id, mode,hash, err := HTML_DOM.GetPayload(); if err != nil {
		return err
	}
	
	if len(origin) | len(id) | len(mode) |  len(hash) == 0 {
		return errors.New("Some of the URLs in your file arent vidmoly matching or are incorrect! See README.md for help")
	}
	
	builder.Log("Info",fmt.Sprintf("Fetch %s Bypass Payload Done! [%s] - %s",*web, hash, ver), "GET Payload", string(typedefs.Purple),"")

	callb(task,&client, web, &origin, &mode, &hash, &id)
	return nil
}
