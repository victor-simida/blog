package main
/*
#include <stdio.h>

void sayHi() {
    printf("Hi");
}
*/
import "C"                /*import C必须独立出来，且对应的注释应当在此行之上*/
import (
	"fmt"
	"runtime"
)

func main() {
	cpus := runtime.GOMAXPROCS(0)            /*cpu数*/
	fmt.Println(cpus)
	C.sayHi()                /*CGO使用*/

}