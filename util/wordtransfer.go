package util

func UpFirst(words string) string {
	if words == "" {
		return ""
	}
	word := []byte (words)
	if word[0] >= 'a' && word[0] <= 'z' {
		word[0] ^= 32
	}
	return string(word)

}

func LowFirst(word string) string {
	if word == "" {
		return ""

	}
	w := []byte(word)
	if w[0] >= 'A' && w[0] <= 'Z' {
		w[0] ^= 32
	}
	return string(w)
}

func HumpToSnake(word string) string {
	if word == "" {
		return ""
	}
	var res []byte
	//var need_ = false
	for idx,ch := range word {
		if idx >0 && ch >='A' && ch <='Z' {
			ch ^=32
			res = append(res, byte('_') )
		}
		res = append(res,byte (ch) )
	}
	return string(res)


}
func SnakeToWord(word string) string {
	var res []byte
	var needUp = false
	for _, ch := range word {
		if  ch =='_' {
			needUp = true
			continue
		}
		if needUp {
			needUp = false
			if ch >='a' && ch <='z' {
				ch ^= 32
			}
		}
		res = append(res, byte (ch) )
	}
	return string(res)
}
