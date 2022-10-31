package algo

/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
*/

/*
哈希表
执行用时：4 ms, 在所有 Go 提交中击败了95.00%的用户
内存消耗：4.1 MB, 在所有 Go 提交中击败了43.01%的用户
通过测试用例：57 / 57
*/
func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range nums {
		loc, ok := hashMap[target-v]
		if ok {
			return []int{i, loc}
		} else {
			hashMap[v] = i
		}
	}
	return nil
}
