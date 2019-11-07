package main

import "fmt"

func addData(a,b int )int {
	return a+b;
}

func main() {
	sum := addData(3,4)
	fmt.Println("两个数的和：",sum)
	Calc()

	//
	fmt.Println("修改了chi/div的数据")
}
