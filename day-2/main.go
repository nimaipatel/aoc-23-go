package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parse_draw(input string) (int, int, int) {
	var red, green, blue int
	for _, val := range strings.Split(input, ",") {
		val = strings.TrimSpace(val)
		parts := strings.Split(val, " ")
		if len(parts) != 2 {
			log.Fatalf("Can't parse %s as a draw", val)
		}
		counts, color := parts[0], parts[1]
		count, err := strconv.Atoi(counts)
		if err != nil {
			log.Fatalf("Can't parse %s as an int", counts)
		}
		switch color {
		case "red":
			red = count
		case "green":
			green = count
		case "blue":
			blue = count
		default:
			log.Fatalf("Can't parse %s as a color", color)
		}
	}
	return red, green, blue
}

func solve_line1(input string) (int, bool) {
	const (
		max_red   = 12
		max_green = 13
		max_blue  = 14
	)

	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		log.Fatalf("Can't parse %q as a game", input)
	}

	var game int
	_, err := fmt.Sscanf(strings.TrimSpace(parts[0]), "Game %d", &game)
	if err != nil {
		log.Fatalf("Can't parse %q as a game", parts[0])
	}

	draws := parts[1]
	for _, draw := range strings.Split(draws, ";") {
		draw = strings.TrimSpace(draw)
		red, green, blue := parse_draw(draw)
		if red > max_red || green > max_green || blue > max_blue {
			return game, false
		}
	}

	return game, true
}

func solve_part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		game, valid := solve_line1(line)
		if valid {
			sum += game
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve_line2(input string) int {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		log.Fatalf("Can't parse %q as a game", input)
	}

	var game int
	_, err := fmt.Sscanf(strings.TrimSpace(parts[0]), "Game %d", &game)
	if err != nil {
		log.Fatalf("Can't parse %q as a game", parts[0])
	}

	draws := parts[1]
	var max_red, max_green, max_blue int
	for _, draw := range strings.Split(draws, ";") {
		draw = strings.TrimSpace(draw)
		red, green, blue := parse_draw(draw)
		max_red = max(red, max_red)
		max_green = max(green, max_green)
		max_blue = max(blue, max_blue)
	}

	return max_red * max_green * max_blue
}

func solve_part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		sum += solve_line2(line)
	}
	return sum
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Please provide an input file")
	}

	file, err := os.Open(args[0])

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	input := string(contents)

	start_part1 := time.Now()
	part1 := solve_part1(input)
	part1_time := time.Since(start_part1)

	start_part2 := time.Now()
	part2 := solve_part2(input)
	part2_time := time.Since(start_part2)

	fmt.Println("Part 1:", part1, "took", part1_time)
	fmt.Println("Part 2:", part2, "took", part2_time)
}
