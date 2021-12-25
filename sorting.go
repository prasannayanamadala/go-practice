package main

import (
	"fmt"
	"sort"
)
func main(){

	fruits := []struct {
		
		Type string
		Quantity int
	}{
		{"Apples-", 12},
		{"Oranges-", 14},
		{"Bananas-", 11},
	}
	sort.SliceStable(fruits, func(i, j int) bool {
		return fruits[i].Quantity > fruits[j].Quantity
	})
	fmt.Println(fruits)

}
