package xpath

import (
	"strings"

	"github.com/antchfx/htmlquery"
)

const (
	payloadInfo = "//td[1]//input[@type='hidden' and @name= 'op' and @value !='login' or @name='hash' or @name='mode' or @name='id']"
)

type Api interface {
	Payload() (origin,id,mode,hash string, err error)
}

type Document struct {
	Htmldoc string
}

// origin, id, mode, hash, err
func (c *Document) GetPayload() (origin,id,mode,hash string, err error) {
	doc, err := htmlquery.Parse(strings.NewReader(c.Htmldoc)); if(err != nil){
		return "","","","",err
	}
	
	data := make(map[string]string)
	inputs := htmlquery.Find(doc, payloadInfo)

	for _ , n := range inputs{
		name := htmlquery.SelectAttr(n,"name")
		value := htmlquery.SelectAttr(n,"value")
	
		data[name] = value
	}

	origin = data["origin"]
	id = data["id"]
	mode = data["mode"]
	hash = data["hash"]

	return
}