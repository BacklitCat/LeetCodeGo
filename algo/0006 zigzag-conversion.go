package algo

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var newStr []byte
	strLen := len(s)
	newStr, loc := make([]byte, strLen), 0
	delta := 2 * (numRows - 1)
	for base := 0; base < numRows; base = base + 1 {
		d := delta - base*2
		i := base
		for i < strLen {
			if d > 0 {
				newStr[loc] = s[i]
				loc++
				i = i + d
			}
			d = delta - d
		}
	}
	return string(newStr)
}
