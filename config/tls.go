package config

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

var (
	TLS_Vers = map[uint16] string{
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
		tls.VersionTLS13: "TLS 1.3",
	}
)


func Config(proxy string) (*http.Transport, error){
	proxyUrl, err := url.Parse(fmt.Sprintf("http://%s",proxy)); if err != nil {
		return nil,err
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
		ForceAttemptHTTP2: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify : true,
		} ,
	}


	return transport,nil
}