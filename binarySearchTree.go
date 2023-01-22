package main

import "fmt"

func main() {
	// 探索対象の配列
	data := []int{1, 3, 7, 13, 17, 21, 74}
	steps, key, _ := linerSearch(data, 21)
	fmt.Println(steps, key)
}

// 線形検索
func linerSearch(data []int, key int) (int, int, error) {
	end := len(data) - 1
	steps := 1

	for i := 0; i < end; i++ {
		if data[i] == key {
			return steps, key, nil
		}
		steps++
	}
	return steps, 0, nil
}
