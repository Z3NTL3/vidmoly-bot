package config

import (
	"crypto/tls"
	"net/http"
)

func Config() (*http.Transport){
	transport := new(http.Transport)
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify : true,
	}

	return transport
}