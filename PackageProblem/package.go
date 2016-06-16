package main
import "fmt"

const N = 3
const W = 50
var value [N + 1][W + 1]int

const (
	VALUE = 0
	WEIGHT = 1
)


/*背包问题*/
var item = [N + 1][2]int{[2]int{0, 0}, [2]int{60, 10}, [2]int{100, 20}, [2]int{120, 30}}
func main() {
	v := pack()
	fmt.Println(v)
}

func pack() int {
	for i := 0; i < W + 1; i ++ {
		value[0][i] = 0
	}

	for i := 0; i < N + 1; i ++ {
		value[i][0] = 0
	}

	for i := 1; i < N + 1; i++ {
		for j := 1; j < W + 1; j++ {
			if j >= item[i][WEIGHT] {
				/*如果添加此物品可以使权重和变大，则添加此物品*/
				if value[i - 1][j - item[i][WEIGHT]] + item[i][VALUE] > value[i - 1][j] {
					value[i][j] = value[i - 1][j - item[i][WEIGHT]] + item[i][VALUE]
				}
			} else {
				value[i][j] = value[i - 1][j]
			}
		}
	}
	return value[N][W]
}



