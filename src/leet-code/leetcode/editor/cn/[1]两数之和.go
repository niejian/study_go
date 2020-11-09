//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。 
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。 
//
// 
//
// 示例: 
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
// 
// Related Topics 数组 哈希表 
// 👍 9559 👎 0
package leetcode

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	//result := []int{-1, -1}
	//// 数组排序
	////sort.Ints(nums)
	//index := -1
	////for i := 0; i < len(nums); i++ {
	////	mytarget := target - nums[i]
	////	index = findTarget(nums, mytarget, i)
	////	if index >= 0 {
	////		result[0] = i
	////		result[1] = index
	////		return result
	////	}
	////}
	//return nil

	return getTargetIndex2(nums, target)

}

//func findTarget(nums []int, target int) int {
//	targetMap := make(map[int] int, 2)
//
//	for index, data := range nums {
//		mytarget := target - nums[index]
//
//		if index > 0 && nums[index] == nums[index - 1] && data != target {
//			continue
//		}
//
//		if data == target && index != targetIndex{
//			return index
//		}
//		targetMap[data] = index
//		if _, ok := targetMap[target]; ok {
//			return target index
//		}
//
//
//	}
//
//	return -1
//}
func getTargetIndex2(nums []int, target int) []int {
	dataMap := make(map[int][]int, 2)
	result := []int{-1, -1}


	for index, data := range nums {
		//numTarget := target - data

		indexs := make([]int, 0)

		// 如果已经存在，跳过
		if _, ok := dataMap[data]; ok {
			indexs = dataMap[data]
			//indexs = append(indexs, index)
			dataMap[data] = append(indexs, index)

		}else {
			dataMap[data] = append(indexs, index)
		}

	}
	fmt.Printf("dataMap: %v \n", dataMap)

	for data, indexs := range dataMap {
		targetData := target - data
		if targetData == data && len(indexs) >= 2 {
			result[0] = indexs[0]
			result[1] = indexs[1]
			return result
		}
		fmt.Printf("indexs: %v \n", indexs)

		// 从map中寻找目标数
		if _, ok := dataMap[targetData]; ok {
			targetIndexs := dataMap[targetData]
			if targetIndexs[0] != indexs[0] && len(targetIndexs) > 0{
				result[0] = indexs[0]
				result[1] = targetIndexs[0]
				return result
			}
		}
	}




	return nil
}

func getTargetIndex(nums []int, target int) []int {
	dataMap := make(map[int]int, 2)
	result := []int{-1, -1}


	for index, data := range nums {
		//numTarget := target - data

		//if data != numTarget && index != 0{
		//	dataMap[data] = index
		//
		//}
		// 如果已经存在，跳过
		if _, ok := dataMap[data]; ok && dataMap[data] == index{
			continue
		}else {
			dataMap[data] = index
		}

	}

	fmt.Printf("dataMap: %v \n", dataMap)

	for data, index := range dataMap {
		numTarget := target - data
		// 是否存在
		if _, ok := dataMap[numTarget]; ok &&  dataMap[numTarget] != index{
			result[0] = index
			result[1] = dataMap[numTarget]
			return result
		}

		if len(dataMap) == 1 && (numTarget == data) {
			result[0] = 0
			result[1] = 1
			return result
		}
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)
