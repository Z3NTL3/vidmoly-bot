package proxies

import (
	"Z3NTL3/Vidmoly-Bot/config"
	"net/http"
	"net/url"
)

func CheckProxies(proxy *string) (error){
	transport := config.Config()
	proxyUrl,err := url.Parse(*proxy); if(err != nil){
		return err
	}
	transport.Proxy = http.ProxyURL(proxyUrl)

	
	return nil
}