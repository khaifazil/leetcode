package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) push(char string) {
	s.items = append(s.items, char)
}

func (s *Stack) pop() {
	s.items = s.items[:len(s.items)-1]
}

func isValid(s string) bool {
	stack := Stack{}

	stack.push(string(s[0]))
	for _, v := range s[1:] {

		l := len(stack.items) - 1

		if len(stack.items) == 0 {
			stack.push(string(v))
		} else {
			switch string(v) {
			case ")":
				if stack.items[l] == "(" {
					stack.pop()
				} else {
					stack.push(string(v))
				}
			case "]":
				if stack.items[l] == "[" {
					stack.pop()
				} else {
					stack.push(string(v))
				}
			case "}":
				if stack.items[l] == "{" {
					stack.pop()
				} else {
					stack.push(string(v))
				}
			default:
				stack.push(string(v))
			}
		}
	}
	if len(stack.items) != 0 {
		return false
	}
	return true
}

func main() {
	str := "(])"

	fmt.Println(isValid(str))
}
