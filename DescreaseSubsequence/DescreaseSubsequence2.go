package main
import "fmt"
/*
	问题描述：求一个数组的最长递减子序列 比如{9，4，3，2，5，4，3，2}的最长递减子序列为{9，5，4，3，2}。
	另外一种解法
*/


func main() {
	origin := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//	origin := []int{1, 35, 43, 2, 3, 455, 1, 3, 6, 3, 46, 3, 4, 7345}
	getSubsequence2(origin)
}

func getSubsequence2(origin []int) int{
	b := make([]int, len(origin))		//生成一个数组b，b[i]标识以origin[i]结尾的最长子序列长度
	b[0] = 1
	for i := 1; i < len(origin); i ++ {
		for j := 0; j < i; j++ {
			if origin[i] < origin[j] && b[i] < b[j] + 1 {
				b[i] = b[j] + 1
			}
		}
	}

	fmt.Println(b)
	var result int
	for _, v := range b {
		if v > result {
			result = v
		}
	}

	fmt.Println(result)
	return result
}

