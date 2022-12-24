package checker

import (
	"regexp"
)

func Website(websiteList []string) (isValid bool) {
	isValid = true
	regex := regexp.MustCompile(`^(?:[a-zA-Z0-9]{1,62}(?:[-\.][a-zA-Z0-9]{1,62})+)(:\d+)?$`)

	for _,val := range websiteList {
		if(!regex.MatchString(val)){
			isValid = false
			break
		}
	}
	return
}