package main

import "fmt"

func main() {
	var arr [4][6]int

	fmt.Println(arr) // [[0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0]]

	// 定义方式
	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1) // [[1 2 3] [4 5 6]]

	var arr2 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2) // [[1 2 3] [4 5 6]]

	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3) // [[1 2 3] [4 5 6]]

	var arr4 = [2][3]string{{"a", "b", "c"}, {"d", "e", "f"}}
	fmt.Println(arr4) // [[a b c] [d e f]]

	arr5 := [2][3]string{{"a", "b", "c"}, {"d", "e", "f"}}
	fmt.Println(arr5) // [[a b c] [d e f]]

	// for 遍历

	for i := 0; i < len(arr5); i++ {
		for j := 0; j < len(arr5[i]); j++ {
			fmt.Printf("%v\t", arr5[i][j])
		}
		fmt.Println()
	}

	// for-range

	for _, v := range arr5 {
		for _, v2 := range v {
			fmt.Printf("%v\t", v2)
		}
		fmt.Println()
	}
}
