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
	times := strings.Fields(lines[0])[1:]
	dists := strings.Fields(lines[1])[1:]

	ways_mult := 1
	for i := 0; i < 4; i++ {
		time, err := strconv.Atoi(times[i])
		check(err)
		dist, err := strconv.Atoi(dists[i])
		check(err)
		ways_mult *= calc_ways(time, dist)
	}

	fmt.Println("Result is:")
	fmt.Println(ways_mult)
}
