//给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。 
//
// 如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。 
//
// 请你返回排序后的数组。 
//
// 
//
// 示例 1： 
//
// 输入：arr = [0,1,2,3,4,5,6,7,8]
//输出：[0,1,2,4,8,3,5,6,7]
//解释：[0] 是唯一一个有 0 个 1 的数。
//[1,2,4,8] 都有 1 个 1 。
//[3,5,6] 有 2 个 1 。
//[7] 有 3 个 1 。
//按照 1 的个数排序得到的结果数组为 [0,1,2,4,8,3,5,6,7]
// 
//
// 示例 2： 
//
// 输入：arr = [1024,512,256,128,64,32,16,8,4,2,1]
//输出：[1,2,4,8,16,32,64,128,256,512,1024]
//解释：数组中所有整数二进制下都只有 1 个 1 ，所以你需要按照数值大小将它们排序。
// 
//
// 示例 3： 
//
// 输入：arr = [10000,10000]
//输出：[10000,10000]
// 
//
// 示例 4： 
//
// 输入：arr = [2,3,5,7,11,13,17,19]
//输出：[2,3,5,17,7,11,13,19]
// 
//
// 示例 5： 
//
// 输入：arr = [10,100,1000,10000]
//输出：[10,100,10000,1000]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr.length <= 500 
// 0 <= arr[i] <= 10^4 
// 
// Related Topics 排序 位运算 
// 👍 74 👎 0
package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func sortByBits(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	// 升序排序
	sort.Ints(arr)
	length := len(arr)
	binArr := make(map[int] []int, length)

	for i := 0; i < length; i++ {
		data := arr[i]
		// 十进制转二进制
		intBin := int2Bin(data)
		//fmt.Printf("%v \n", intBin )
		// 统计1的个数
		count := strings.Count(intBin, "1")
		//fmt.Printf("data: %v, 二进制：%v，1的数量 %v \n", data, intBin, count )
		if _, ok := binArr[count]; ok {
			countArr := binArr[count]
			//sort.Ints(countArr)
			countArr = append(countArr, data)
			binArr[count] = countArr
			//fmt.Printf("count: %v exist: %v , countArr: %v \n", count, ok, countArr)

			// 排序
			sort.Ints(countArr)
		}else {
			countArr := make([]int, 0)
			binArr[count] = append(countArr, data)
		}
		//if _, ok := binArr[count]; !ok {
		//	countArr := make([]int, 0)
		//	binArr[count] = append(countArr, data)
		//}


	}
	//fmt.Printf("binArr--> %v", binArr )


	result := make([]int, 0)
	keys := make([]int, 0)
	for key, _ := range binArr {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	if len(keys) > 0 {
		for _, key := range keys {
			datas := binArr[key]
			for _, i := range datas {
				result = append(result, i)
			}
		}
	}

	return result
}



// 十进制转二进制
func int2Bin(data int) string {
	result := ""

	if data == 0 {
		return "0"
	}

	for ;data > 0;data /= 2 {
		lsb := data % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)
