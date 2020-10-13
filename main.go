package main

import "fmt"

func main()  {
	var s []int
	var n int
	var sum int
	var num int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		s = append(s, num)
		sum += num
	}

	fmt.Println(sum)
}
