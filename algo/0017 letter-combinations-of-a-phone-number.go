package algo

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var combinations []string
	var backtrack func(string, int, string)
	backtrack = func(digits string, index int, combination string) {
		if index == len(digits) {
			combinations = append(combinations, combination)
		} else {
			digit := string(digits[index])
			letters := phoneMap[digit]
			lettersCount := len(letters)
			for i := 0; i < lettersCount; i++ {
				backtrack(digits, index+1, combination+string(letters[i]))
			}
		}
	}

	backtrack(digits, 0, "")
	return combinations
}

/*
var phoneMap map[string]string = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}

var finalResult []string
var inputDigits string

func letterCombinations(digits string) []string {
    if len(digits)==0{
        return []string{}
    }
    inputDigits = digits
    backtrack(0,"")
    return finalResult
}

func backtrack(index int, tempSearch string){
    if index == len(inputDigits){
        finalResult = append(finalResult, tempSearch)
        return
    }
    appendStr := phoneMap[string(inputDigits[index])]
    for i:=0;i<len(appendStr);i++{
        backtrack(index+1, tempSearch+string(appendStr[i]))
    }
}
*/
