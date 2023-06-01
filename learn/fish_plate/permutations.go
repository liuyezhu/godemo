package fush_plate

import "fmt"

func permutation(numbers []int, result [][]int, temp []int) [][]int {
	if len(temp) == len(numbers) {
		result = append(result, temp)
		return result
	}
	for i := 0; i < len(numbers); i++ {
		if !contains(temp, numbers[i]) {
			newTemp := make([]int, len(temp))
			copy(newTemp, temp)
			newTemp = append(newTemp, numbers[i])
			result = permutation(numbers, result, newTemp)
		}
	}
	return result
}

func combination(numbers []int, result [][]int, temp []int, start int, k int) [][]int {
	if len(temp) == k {
		newTemp := make([]int, len(temp))
		copy(newTemp, temp)
		result = append(result, newTemp)
		return result
	}
	for i := start; i < len(numbers); i++ {
		temp = append(temp, numbers[i])
		result = combination(numbers, result, temp, i+1, k)
		temp = temp[:len(temp)-1]
	}
	return result
}

func contains(numbers []int, target int) bool {
	for _, n := range numbers {
		if n == target {
			return true
		}
	}
	return false
}

func PermAction() {
	numbers := []int{1, 2, 3, 4, 5}
	// 生成所有排列
	permutations := [][]int{}
	permutations = permutation(numbers, permutations, []int{})
	fmt.Println("All permutations:", permutations)

	// 生成所有组合
	combinations := [][]int{}
	for k := 1; k <= len(numbers); k++ {
		combinations = combination(numbers, combinations, []int{}, 0, k)
	}
	fmt.Println("All combinations:", combinations)
}
