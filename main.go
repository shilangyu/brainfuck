package main

import (
	"flag"
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

const cellAmount = 30000

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	c := flag.String("c", "", "brainfuck code to be interpreted")
	flag.Parse()
	if *c == "" {
		file := os.Args[1]
		content, _ := ioutil.ReadFile(file)
		*c = string(content)
	}
	code := *c

	var memory [cellAmount]uint8
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
					pointer = cellAmount - 1
				} else {
					pointer--
				}
			}
		case '>':
			if skippingLoop == 0 {
				if pointer == cellAmount-1 {
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
