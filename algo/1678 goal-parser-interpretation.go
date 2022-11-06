package algo

import "strings"

func interpretSb(command string) string {
	var sb strings.Builder
	cmdLen := len(command)
	for i := 0; i < cmdLen; i++ {
		if command[i] == 'G' {
			sb.WriteByte('G')
		} else if command[i] == ')' && command[i-1] == '(' {
			sb.WriteByte('o')
		} else if command[i] == ')' && command[i-1] == 'l' {
			sb.Write([]byte("al"))
		}
	}
	return sb.String()
}

func interpret(command string) string {
	var res []byte
	cmdLen := len(command)
	for i := 0; i < cmdLen; i++ {
		if command[i] == 'G' {
			res = append(res, 'G')
		} else if command[i] == ')' && command[i-1] == '(' {
			res = append(res, 'o')
		} else if command[i] == ')' && command[i-1] == 'l' {
			res = append(res, 'a')
			res = append(res, 'l')
		}
	}
	return string(res)
}
