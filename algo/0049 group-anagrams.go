package algo

import "log"

func Test() {
	log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	var (
		locateMap map[[26]int][]string = make(map[[26]int][]string, 1)
		res       [][]string
	)
	for i := 0; i < len(strs); i++ {
		k := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			k[strs[i][j]-97]++
		}
		if _, ok := locateMap[k]; ok {
			locateMap[k] = append(locateMap[k], strs[i])
		} else {
			locateMap[k] = []string{strs[i]}
		}
	}
	for _, v := range locateMap {
		res = append(res, v)
	}
	return res
}
