package main

import "fmt"

func main()  {
	fmt.Println("Hello")
	for i := 0; i < 10; i ++ {
		fmt.Println(i)
	}
	
	if 1 == 2 {
		fmt.Println("1==1")
	} else {
		fmt.Println("1 != 2")
	}
}