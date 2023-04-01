package pkg

func Chekport(port string) bool {
	runes := []rune(port)
	for _, v := range runes {
		if (v >= '0' && '9' >= v) && (len(port) == 4) {
			return true
		}
	}
	return false
}