package bot

import (
	"Z3NTL3/Vidmoly-Bot/builder"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"Z3NTL3/Vidmoly-Bot/xpath"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// // origin, id, mode, hash, err
func Bypass(
	taskid int,
	client *http.Client,
	uri *string, 
	origin *string,
	mode *string,
	hash *string,
	id *string,
) {
	rand.Seed(time.Now().Unix())
	payload,err := url.ParseQuery(fmt.Sprintf("op=%s&id=%s&mode=%s&hash=%s", *origin, *id, *mode, *hash)); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}

	fmt.Println(payload)

	req, err := http.NewRequest("POST", *uri, strings.NewReader(payload.Encode())); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", typedefs.Headers_[rand.Intn(len(typedefs.Headers_))])

	resp, err := client.Do(req); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}
	bodyFull, err := io.ReadAll(resp.Body); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}
	defer resp.Body.Close()

	HTML_DOM := xpath.Document{Htmldoc: string(bodyFull)}
	src, err := HTML_DOM.GetDownloadSource(); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}
	// komt nog check en file saver
	builder.Log("Task",fmt.Sprintf("Task %d completed - %s Saved to save dir file ", taskid, *src) ,"Task", string(typedefs.Purple),"")

}