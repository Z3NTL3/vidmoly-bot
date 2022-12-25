package fetch

import (
	"Z3NTL3/Vidmoly-Bot/checker"
	"Z3NTL3/Vidmoly-Bot/config"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"Z3NTL3/Vidmoly-Bot/xpath"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)
func InitBypass(web, proxy *string, callb typedefs.BypassType) (error){
	err := checker.CheckProxy(proxy); if(err != nil){
		return err
	}

	transport := config.Config()
	client := http.Client{
		Transport: transport,
	}
	var bodyReader io.ReadCloser
	req, err := http.NewRequestWithContext(context.Background(), "GET", *web, bodyReader); if(err != nil){
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
	bodyFull := string(body[0:length])
	_ = config.TLS_Vers[resp.TLS.Version]

	HTML_DOM := xpath.Document{Htmldoc: bodyFull}
	origin, id, mode,hash, err := HTML_DOM.GetPayload(); if err != nil {
		return err
	}
	
	if len(origin) | len(id) | len(mode) |  len(hash) == 0 {
		return errors.New("Some of the URLs in your file arent vidmoly matching or are incorrect! See README.md for help")
	}

	// bypass stuk nog te doen
	return nil
}