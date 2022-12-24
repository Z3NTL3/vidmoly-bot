package config

import (
	"crypto/tls"
	"net/http"
)

var (
	TLS_Vers = map[uint16] string{
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
		tls.VersionTLS13: "TLS 1.3",
	}
)


func Config() (*http.Transport){
	transport := new(http.Transport)
	
	transport.ForceAttemptHTTP2 = true
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify : false,
	}

	return transport
}