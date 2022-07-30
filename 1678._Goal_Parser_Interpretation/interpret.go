package main

import "fmt"

func main() {
	cmd := "(al)G(al)()()G"

	fmt.Println(interpret(cmd))
}
func interpret(command string) (res string) {

	ptr1, ptr2 := 0, 1
	for ptr1 < len(command) {
		if command[ptr1] == 'G' {
			res += "G"
			ptr1++
			ptr2++
		} else if command[ptr1] == '(' && command[ptr2] == ')' {
			res += "o"
			ptr1 += 2
			ptr2 += 2
		} else if command[ptr1] == '(' && command[ptr2] == 'a' {
			res += "al"
			ptr1 += 4
			ptr2 += 4
		}

	}
	return
}
