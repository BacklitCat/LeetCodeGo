package algo

import (
	"container/list"
	"fmt"
	"strings"
)

func test() {
	fmt.Println(parseBoolExpr("&(t,t,|(t,f,&(f,f),t),t)")) // true
	fmt.Println(parseBoolExpr("!(f)"))                     // true
	fmt.Println(parseBoolExpr("|(f,t)"))                   // true
	fmt.Println(parseBoolExpr("&(t,f)"))                   // false
	fmt.Println(parseBoolExpr("|(&(t,f,t),!(t))"))         // false
}

var parseMap map[byte]func(expression string) bool = map[byte]func(expression string) bool{
	'&': parseAnd,
	'|': parseOr,
	'!': parseNot,
}

var bool2char map[bool]byte = map[bool]byte{
	true:  't',
	false: 'f',
}

func parseBoolExpr(expression string) bool {
	stack := list.New()
	expLen := len(expression)
	var newBool bool
	for i := 0; i < expLen; i++ {
		if expression[i] != ')' {
			if expression[i] != ',' {
				stack.PushBack(expression[i])
			}
		} else {
			var sb strings.Builder
			c := pop(stack)
			for c != '(' {
				sb.WriteByte(c)
				c = pop(stack)
			}
			c = pop(stack) // & | !
			newBool = parseMap[c](sb.String())
			stack.PushBack(bool2char[newBool])
		}
	}
	return newBool

}

func pop(stack *list.List) byte {
	e := stack.Back()
	if e != nil {
		stack.Remove(e)
		return e.Value.(byte)
	}
	return 0
}

func parseOr(expression string) bool {
	expLen := len(expression)
	for i := 0; i < expLen; i++ {
		if expression[i] == 't' {
			return true
		}
	}
	return false
}

func parseAnd(expression string) bool {
	expLen := len(expression)
	for i := 0; i < expLen; i++ {
		if expression[i] == 'f' {
			return false
		}
	}
	return true
}

func parseNot(expression string) bool {
	if expression[0] == 'f' {
		return true
	}
	return false
}
