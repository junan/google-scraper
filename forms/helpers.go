package forms

import "unicode"

func firstCapitalise(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])

	return string(tmp)
}
