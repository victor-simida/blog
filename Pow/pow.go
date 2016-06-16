package main
import (
	"time"
	"fmt"
	"strings"
)

/*******************************************
*函数名：pow1
*作用：求sq的times次幂
*时间：2016/6/12 16:26
*******************************************/
func pow1(sq, times int) int {
	pow := 1
	for i := times; i > 0; i >>= 1 {
		if i & 1 != 0 {
			pow *= sq
		}
		sq *= sq
	}

	return pow
}

/*******************************************
*函数名：pow2
*作用：求sq的times次幂
*时间：2016/6/12 16:26
*******************************************/
func pow2(sq, times int) int {
	pow := 1
	for i := times; i > 0; i -= 1 {
		pow *= sq
	}
	return pow
}

func main() {
	var rst int
	s1 := time.Now()
	for i := 1; i <= 200000; i ++ {
		rst = pow1(999, 1000000000000)
	}
	e1 := time.Now()
	fmt.Println(e1.Sub(s1).Nanoseconds())
	fmt.Printf("Pow1 Resturn %v\n", rst)


	s2 := time.Now()
	for i := 1; i <= 200000; i ++ {
		rst = pow2(999, 1000000000000)
	}
	e2 := time.Now()
	fmt.Println(e2.Sub(s2).Nanoseconds())
	fmt.Printf("Pow2 Resturn %v\n", rst)
}


/*
运行结果:
7000400
Pow1 Resturn 4012420759885938689
16001000
Pow2 Resturn -6165855492271999103


可以看到前者性能上面的优势非常明显

*/


