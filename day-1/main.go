package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func parse_digit(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	} else {
		return -1
	}
}

func parse_word_digit(input string) int {
	if strings.HasPrefix(input, "zero") {
		return 0
	} else if strings.HasPrefix(input, "one") {
		return 1
	} else if strings.HasPrefix(input, "two") {
		return 2
	} else if strings.HasPrefix(input, "three") {
		return 3
	} else if strings.HasPrefix(input, "four") {
		return 4
	} else if strings.HasPrefix(input, "five") {
		return 5
	} else if strings.HasPrefix(input, "six") {
		return 6
	} else if strings.HasPrefix(input, "seven") {
		return 7
	} else if strings.HasPrefix(input, "eight") {
		return 8
	} else if strings.HasPrefix(input, "nine") {
		return 9
	} else {
		return -1
	}
}

func update(digit int, fst, lst *int) {
	if *fst == -1 {
		*fst = digit
	}
	*lst = digit
}

func solve_part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		fst, lst := -1, -1
		for _, char := range line {
			if digit := parse_digit(char); digit >= 0 {
				update(digit, &fst, &lst)
			}
		}
		if fst != -1 && lst != -1 {
			sum += fst*10 + lst
		}
	}
	return sum
}

func solve_part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		fst, lst := -1, -1
		for i, char := range line {
			if digit := parse_digit(char); digit >= 0 {
				update(int(char-'0'), &fst, &lst)
			} else if word_digit := parse_word_digit(line[i:]); word_digit >= 0 {
				update(word_digit, &fst, &lst)
			}
		}
		if fst != -1 && lst != -1 {
			sum += fst*10 + lst
		}
	}
	return sum
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Expected input file name as only argument")
	}

	file, err := os.Open(args[0])
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	file_contents, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	input := string(file_contents)

	start_part1 := time.Now()
	part1 := solve_part1(input)
	part1_time := time.Since(start_part1)

	start_part2 := time.Now()
	part2 := solve_part2(input)
	part2_time := time.Since(start_part2)

	fmt.Println("Part 1: ", part1, "took ", part1_time)
	fmt.Println("Part 2: ", part2, "took ", part2_time)
}
