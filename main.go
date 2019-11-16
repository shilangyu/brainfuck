package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type peekStack []int

func (s *peekStack) Push(item int) {
	*s = append(*s, item)
}

func (s peekStack) Peek() int {
	return s[len(s)-1]
}

func (s *peekStack) Pop() int {
	val := s.Peek()
	*s = (*s)[:len(*s)-1]
	return val
}

func main() {
	file := os.Args[1]
	content, _ := ioutil.ReadFile(file)
	code := string(content)

	var memory [30000]uint8
	var pointer uint
	var loops peekStack
	var skippingLoop int

	for i := 0; i < len(code); i++ {
		char := code[i]

		switch char {
		case '-':
			if skippingLoop == 0 {
				memory[pointer]--
			}
		case '+':
			if skippingLoop == 0 {
				memory[pointer]++
			}
		case '<':
			if skippingLoop == 0 {
				if pointer == 0 {
					pointer = 29999
				} else {
					pointer--
				}
			}
		case '>':
			if skippingLoop == 0 {
				if pointer == 29999 {
					pointer = 0
				} else {
					pointer++
				}
			}
		case '[':
			if skippingLoop == 0 {
				if memory[pointer] != 0 {
					loops.Push(i)
				} else {
					skippingLoop++
				}
			} else {
				skippingLoop++
			}
		case ']':
			if skippingLoop > 0 {
				skippingLoop--
			} else {
				if memory[pointer] != 0 {
					i = loops.Peek()
				} else {
					loops.Pop()
				}
			}
		case ',':
			if skippingLoop == 0 {
				fmt.Scanf("%c", &memory[pointer])
			}
		case '.':
			if skippingLoop == 0 {
				fmt.Print(string(rune(memory[pointer])))
			}
		}
	}

}
