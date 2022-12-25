package checker

func CheckInfo(m map[string][]string) (valid bool) {
	valid = true
	if len(m["links"]) == 0 || len(m["proxies"]) == 0 {
		valid = false
	}
	return
}