package checker

import (
	"Z3NTL3/Vidmoly-Bot/config"
	"net/http"
	"net/url"
)

func CheckProxy(proxy *string) error {
	transport := config.Config()
	proxyUrl, err := url.Parse(*proxy)
	if err != nil {
		return err
	}
	transport.Proxy = http.ProxyURL(proxyUrl)

	client := new(http.Client)
	resp, err := client.Get("https://pix4.dev")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	// error will throw if it cannot connect to the proxy

	return nil
}