package main

import "fmt"

// 遍历数组，在map中寻找另一半数字，如果找到了，直接返回2个数字的下标即可
// 如果找不到，就把这个数字存入map中
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// 例如nums=[2,7,11,15], target=9，寻找的数字是9-2=7
	for k, v := range nums {
		//判断m[9-2]是否存在
		if idx, ok := m[target-v]; ok {
			// 存在，返回[0，1]
			return []int{idx, k}
		}
		// 不存在，将m[2]=0
		m[v] = k
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	// 输入：nums = [2,7,11,15], target = 9
	// 输出：[0,1]
	fmt.Printf("nums = %v, target = %v\n", nums, target)
	fmt.Printf("output = %v\n", twoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Printf("nums = %v, target = %v\n", nums, target)
	fmt.Printf("output = %v\n", twoSum(nums, target))
}
