package bot

import (
	"Z3NTL3/Vidmoly-Bot/builder"
	"Z3NTL3/Vidmoly-Bot/typedefs"
	"Z3NTL3/Vidmoly-Bot/xpath"
	"compress/gzip"
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

	req, err := http.NewRequest("POST", *uri, strings.NewReader(payload.Encode())); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", typedefs.Headers_[rand.Intn(len(typedefs.Headers_))])
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Keep-Alive", "max=650000")

	resp, err := client.Do(req); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}
	gzip, err := gzip.NewReader(resp.Body); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}

	bodyFull, err := io.ReadAll(gzip); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}
	defer resp.Body.Close()

	HTML_DOM := xpath.Document{Htmldoc: string(bodyFull)}
	src, err := HTML_DOM.GetDownloadSource(); if err != nil {
		builder.Log("Err Info",err.Error(), "Err", string(typedefs.Red),"\n")
		return
	}
	
	if len(*src) != 0 {
		builder.Log("Task",fmt.Sprintf("Task %d completed - %s Saved to save dir file ", taskid, *src) ,"Task", string(typedefs.Purple),"")
	} else {
		builder.Log("Task",fmt.Sprintf("Task %d Failed - %s Didnt bypass... For unknown reason ", taskid, *src) ,"Task", string(typedefs.Red),"")
	}

	fmt.Println(resp.Request.Header)

}