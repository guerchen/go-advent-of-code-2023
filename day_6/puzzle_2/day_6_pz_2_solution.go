package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calc_ways(time int, dist int) int {
	ways := 0
	for i := 0; i <= time; i++ {
		moving_time := time - i
		moved_dist := i * moving_time
		if moved_dist > dist {
			ways += 1
		}
	}
	return ways
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	input := string(dat)
	lines := strings.Split(input, "\n")
	time := strings.Join(strings.Fields(lines[0])[1:], "")
	dist := strings.Join(strings.Fields(lines[1])[1:], "")

	_time, err := strconv.Atoi(time)
	check(err)
	_dist, err := strconv.Atoi(dist)
	check(err)
	ways := calc_ways(_time, _dist)

	fmt.Println("Result is:")
	fmt.Println(ways)
}
