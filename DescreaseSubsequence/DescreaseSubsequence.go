package main
import "fmt"

/*
	问题描述：求一个数组的最长递减子序列 比如{9，4，3，2，5，4，3，2}的最长递减子序列为{9，5，4，3，2}。
*/

var LastResult [][]int
func main() {
	LastResult = make([][]int, 0)
	origin := []int{9,8,7,6,5,4,3,2,1}
//	origin := []int{1, 35, 43, 2, 3, 455, 1, 3, 6, 3, 46, 3, 4, 7345}
	var result []int
	getSubsequence(origin, result, 0)

	fmt.Println(LastResult)

	max := 0
	last := []int{}
	for _,v := range LastResult {
		if len(v) > max {
			last = v
			max = len(v)
		}
	}

	fmt.Println(max)
	fmt.Println(last)
}

func getSubsequence(origin, result []int, index int) {
	if index >= len(origin)  {
		if len(result) != 0 {
			LastResult = append(LastResult, result)
		}
		return
	}

	/*如果目标数大于结果数组中最后一个数字，则跳过*/
	already := len(result)
	if already != 0 && result[already - 1] <= origin[index] {
		getSubsequence(origin, result, index + 1)
		return
	}

	clone1 := append(result, origin[index])
	clone2 := append([]int{}, result...)
	getSubsequence(origin, clone1, index + 1)
	getSubsequence(origin, clone2, index + 1)
}

