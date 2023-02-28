package algo

import "log"

func test139() {
	log.Println(wordBreak("a", []string{"a"}))
	log.Println(wordBreak("a", []string{"b"}))
	log.Println(wordBreak("a", []string{"a", "b"}))
	log.Println(wordBreak("aaaab", []string{"a", "aa"}))
	log.Println(wordBreak("leetcode", []string{"leet", "code"}))
	log.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	log.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	log.Println(wordBreak("abcd", []string{"a", "abc", "b", "cd"})) //true
	log.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"})) //true
}

func wordBreak(s string, wordDict []string) bool {
	var (
		res   bool
		stop  bool
		dfs   func(left int)
		cache = make([]bool, len(s)+1)
	)
	dfs = func(left int) {
		if stop || left > len(s) {
			return
		}
		if left == len(s) {
			res = true
			stop = true
			return
		}
		for i := 0; i < len(wordDict); i++ {
			l := len(wordDict[i])
			if left+l <= len(s) && s[left:left+l] == wordDict[i] {
				if cache[left+l] {
					left = left + l
					i = 0
				} else {
					cache[left+l] = true
					dfs(left + l)
					//log.Println("dfs(left + l)", left, l, left+l)
				}
			}
		}

	}
	dfs(0)
	return res
}
