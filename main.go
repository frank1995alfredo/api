package main

import (
	"fmt"

	_ "github.com/frank1995alfredo/api/routes"
)

func main() {

	//routes.Rutas()

	var nums []int
	nums = make([]int, 3)
	nums[0] = 5
	fmt.Println(nums)

}
