package checker

import (
	"Z3NTL3/Vidmoly-Bot/builder"
	"Z3NTL3/Vidmoly-Bot/config"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"fmt"
	"net/http"
)

func CheckProxy(proxy *string) error {
	transport,err := config.Config(*proxy); if err != nil {
		return err
	}

	builder.Log("CHECK", fmt.Sprintf("Checking proxy: %s", *proxy),"Checking", string(typedefs.Purple),"")

	client := &http.Client{Transport: transport}
	resp, err := client.Get("https://pix4.dev")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	builder.Log("CHECK", fmt.Sprintf("Proxy Valid!: %s", *proxy),"Valid", string(typedefs.LightPurple),"")
	// error will throw if it cannot connect to the proxy

	return nil
}