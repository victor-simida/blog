package main
import "fmt"


var x []int = []int{34, 6, 7, 3, 345, 23, 345, 745, 234, 6, 3, 6, 73, 6}

/*非常简单的快速排序*/
func quickSort(low, high int) {
	if low >= high {
		return
	}

	m := low
	key := x[m]
	for i := low + 1; i < high; i ++ {
		if x[i] < key {
			m ++
			swap(m, i)
		}
	}

	swap(low, m)
	quickSort(low, m)
	quickSort(m + 1, high)

}
func main() {
	fmt.Println("Before:")
	fmt.Println(x)

	quickSort(0,len(x))

	fmt.Println("After:")
	fmt.Println(x)

}

func swap(a, b int) {
	temp := x[a]
	x[a] = x[b]
	x[b] = temp
}
