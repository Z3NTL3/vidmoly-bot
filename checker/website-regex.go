package checker

import (
	"regexp"
)

func Website(websiteList []string) (isValid bool) {
	isValid = true
	regex := regexp.MustCompile(`^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\?([^#]*))?(#(.*))?`)

	for _,val := range websiteList {
		// fmt.Println(val,regex.MatchString(val))
		
		if(!regex.MatchString(val)){
			isValid = false
			break
		}
	}
	return
}