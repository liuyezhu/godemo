package fush_plate

import (
	"math/rand"
	"time"
)

func GenerateCards(t int) [][]int {

	rand.Seed(time.Now().UnixNano())
	cards := make([]int, 54)
	for i := 1; i <= 54; i++ {
		cards[i-1] = i
	}

	var cardsGroups [][]int
	for j := 1; j <= t; j++ {
		cardsGroup := make([]int, 4)
		for k := 1; k <= 4; k++ {
			index := rand.Intn(len(cards) - 1)
			cardsGroup[k-1] = cards[index]
			cards[index] = cards[len(cards)-1]
			cards = cards[:len(cards)-1]
		}
		cardsGroups = append(cardsGroups, cardsGroup)
	}

	return cardsGroups
}

func SortCardSize(cards [][]int) {

}

func AutoCardPicking(cards []int) {

	//if len(ArrayIntersect([]int{1, 21, 3}, []int{10, 323})) > 0 {
	//
	//}
	//
	//s := make([][2]int, 4)
	//for k, v := range cards {
	//	points := v % 13
	//	if inSlice(points, []int{10, 11, 12}) {
	//		v = 0
	//	}
	//	s[k] = [2]int{v, points}
	//}
	//
	//sort.Slice(s, func(i, j int) bool {
	//	return s[i][1] < s[j][1]
	//})

	//fmt.Println(s)
}

func inSlice(val int, slice []int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ArrayIntersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	var res []int
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			res = append(res, n)
			m[n] = false
		}
	}
	return res
}
