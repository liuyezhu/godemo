package fush_plate

import (
	"fmt"
	"math/rand"
)

func main() {
	s := []string{"apple", "banana", "cherry", "date", "elderberry", "fig"}
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	fmt.Println(s)
}
