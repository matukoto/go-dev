package main

import "fmt"

func main() {
	// 探索対象の配列
	data := []int{1, 3, 7, 13, 17, 21, 74}
	steps, key, _ := linerSearch(data, 21)
	fmt.Println(steps, key)

	search("線形探索", data, linerSearch)
}

// 線形検索
func linerSearch(data []int, key int) (int, int, error) {
	end := len(data)
	steps := 1

	for i := 0; i < end; i++ {
		if data[i] == key {
			return steps, key, nil
		}
		steps++
	}
	return steps, 0, nil
}

// 関数の実行
// 最大ステップ数と平均ステップ数を求める
func search(title string, data []int, fn func(list []int, value int) (int, int, error)) {
	fmt.Println(title)
	maxSteps := 1
	stepSum := 0
	for i, v := range data {
		steps, key, _ := fn(data, v)
		stepSum += steps
		if maxSteps < steps {
			maxSteps = steps
		}
		fmt.Println("key: ", key, "position: ", i, "steps: ", steps)
	}
	fmt.Println("最大ステップ数: ", maxSteps)
	fmt.Println("平均ステップ数: ", stepSum/len(data))
}
