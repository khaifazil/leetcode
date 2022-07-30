package main

import "fmt"

func main() {
	order := "abcdefghijklmnopqrstuvwxyz"
	words := []string{"apple", "app"}

	fmt.Println(isAlienSorted(words, order))

}

func isAlienSorted(words []string, order string) bool {
	ptr1, ptr2 := 0, 1
	m := mapOrder(order)
	isOrdered := true
	for isOrdered && ptr2 < len(words) {
		for i := 0; i < len(words[ptr1]); i++ {
			strA := words[ptr1]
			strB := words[ptr2]
			if i > len(strB)-1 {
				isOrdered = false
				break
			}
			if m[int32(strA[i])] < m[int32(strB[i])] {
				break
			} else if m[int32(strA[i])] > m[int32(strB[i])] {
				isOrdered = false
				break
			}
		}
		ptr1++
		ptr2++
	}

	return isOrdered
}

func mapOrder(order string) map[int32]int {
	m := make(map[int32]int)

	for i, v := range order {
		m[v] = i + 1
	}
	return m
}
