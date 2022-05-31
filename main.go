package main

import (
	"fmt"
	"protek/duplicate"
)

func main() {
	arr := [][]string{
		{"1014065", "1", "4708", "4709"},
		{"1014071", "1", "4708", "4697"},
		{"1014130", "2", "9416", "8004"},
		{"1014130", "2", "9416", "1412"},
		{"1014238", "1", "4708", "4694"},
		{"1014240", "3", "14124", "14090"},
	}

	result := duplicate.Find(arr)

	for _, data := range result {
		fmt.Println(data)
	}
}
