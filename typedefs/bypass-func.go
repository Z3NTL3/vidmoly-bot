package typedefs

import "net/http"


type BypassType func(
	taskid int,
	client *http.Client,
	uri *string, 
	origin *string,
	mode *string,
	hash *string,
	id *string,
)