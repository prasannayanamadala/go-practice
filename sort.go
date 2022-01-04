package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("sorting fruits:\n")
	
	str := []string{"bananas\n","oranges\n","apples\n"}
	sort.Strings(str)
	fmt.Println("\n",str)
	
	int := []int{14,11,12}
	sort.Ints(int)
	fmt.Println("\n",int)
}



