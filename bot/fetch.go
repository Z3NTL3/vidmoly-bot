package bot

import (
	"Z3NTL3/Vidmoly-Bot/config"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"
)
func Fetch(webList, proxies *[]string, callb func()){
	transport := config.Config()
	client := http.Client{
		Transport: transport,
	}
	bodyReader := new(io.Reader)
	
	req, err := http.NewRequestWithContext(context.Background(), "GET", "https://pix4.dev", *bodyReader); if(err != nil){
		fmt.Println(err)
	}

	resp, err := client.Do(req); if(err != nil){
		fmt.Println(err)
	}

	var body string

	for {
		buffer := make([]byte,1042)
		len, err := resp.Body.Read(buffer); if(err != nil){
			break
		}

		body += string(buffer[0:len])
	}
	tls_maps := map[uint16]string{
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
		tls.VersionTLS13: "TLS 1.3",
	}


	fmt.Println(body)
	fmt.Println(tls_maps[resp.TLS.Version])
	
}