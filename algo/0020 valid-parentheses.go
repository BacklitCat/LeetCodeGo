package algo

import (
	"container/list"
)

func isValid(s string) bool {
	l := list.New()
	serLen := len(s)
	for i := 0; i < serLen; i++ {
		backItem := l.Back()
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			l.PushBack(s[i])
		} else if backItem == nil {
			return false
		} else if s[i] == ']' {
			if backItem.Value.(uint8) == '[' {
				l.Remove(l.Back())
			} else {
				return false
			}
		} else if s[i] == ')' {
			if backItem.Value.(uint8) == '(' {
				l.Remove(l.Back())
			} else {
				return false
			}
		} else if s[i] == '}' {
			if backItem.Value.(uint8) == '{' {
				l.Remove(l.Back())
			} else {
				return false
			}
		}
	}
	if l.Front() == nil {
		return true
	}
	return false
}

func isValidAnswer(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
