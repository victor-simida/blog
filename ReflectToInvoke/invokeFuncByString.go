package main

import (
	"fmt"
	"reflect"
)

var opFuncs = map[string]string{"+": "Add",
	"-": "Minus",
	"*": "Multi",
	"/": "Divide",
}

type Calculator struct {
}

func (this *Calculator) Add(a, b int) int {
	return a + b
}

func (this *Calculator) Minus(a, b int) int {
	return a - b
}

func (this *Calculator) Multi(a, b int) int {
	return a * b
}

func (this *Calculator) Divide(a, b int) int {
	return a / b
}

func main() {
	var a int = 10
	var b int = 20

	calc := &Calculator{}

	for funcOp, funcName := range opFuncs {
		fn, _ := reflect.TypeOf(calc).MethodByName(funcName)

		params := make([]reflect.Value, 3)

		params[0] = reflect.ValueOf(calc)		//由于是类型方法，第一个param应当是类型本身
		params[1] = reflect.ValueOf(a)
		params[2] = reflect.ValueOf(b)

		v := fn.Func.Call(params)
		fmt.Println(fmt.Sprintf("%d %s %d = %d", a, funcOp, b, v[0].Int()))
	}
}
