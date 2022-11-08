package algo

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	var allowedHashSet map[byte]bool = map[byte]bool{}
	for i := 0; i < len(allowed); i++ {
		allowedHashSet[allowed[i]] = true
	}
	count := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			_, ok := allowedHashSet[words[i][j]]
			if !ok {
				break
			} else if j == len(words[i])-1 {
				fmt.Println(words[i], "is Consistent String")
				count++
			}
		}
	}
	return count
}
