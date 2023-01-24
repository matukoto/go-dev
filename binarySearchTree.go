package main

import "fmt"
import "errors"

func main() {
	// 探索対象の配列
	data := []int{1, 3, 7, 18, 17, 21, 74}

	search("線形探索", data, linerSearch)
	search("二分探索", data, binarySearch)
}

// 線形検索
func linerSearch(data []int, key int) (int, int, error) {
	end := len(data)
	steps := 0

	for i := 0; i < end; i++ {
		steps++
		if data[i] == key {
			return steps, key, nil
		}
	}
	return 0, 0, errors.New("Key not found")
}

// 二分探索木
func binarySearch(data []int, key int) (int, int, error) {

	// 探索対象の配列がソートされていることを確認
	for i := 0; i < len(data)-1; i++ {

		if data[i] < data[i+1] {
			return 0, 0, errors.New("Array is not sorted")
		}
	}

	start := 0
	end := len(data) - 1
	steps := 0

	for start <= end {

		// オーバーフローの可能性があるため
		middle := start + (end-start)/2
		steps++

		if data[middle] == key {
			return steps, key, nil
		} else if data[middle] < key {
			start = middle + 1
		} else {
			end = middle - 1
		}

	}

	return steps, key, errors.New("Key not found")
}

// 関数の実行
// 最大ステップ数と平均ステップ数を求める
func search(title string, data []int, fn func(list []int, value int) (int, int, error)) {
	fmt.Println(title)
	maxSteps := 1
	stepSum := 0

	for i, v := range data {
		steps, key, err := fn(data, v)
		if err != nil {
			fmt.Println(err)
			return
		}
		stepSum += steps

		if maxSteps < steps {
			maxSteps = steps
		}

		fmt.Println("key: ", key, "position: ", i, "steps: ", steps)
	}

	fmt.Println("最大ステップ数: ", maxSteps)
	fmt.Println("平均ステップ数: ", stepSum/len(data))
}
