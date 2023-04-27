package util

func StringsContainString(strings []*string, str string) bool {
	for _, value := range strings {
		if *value == str {
			return true
		}
	}
	return false
}
