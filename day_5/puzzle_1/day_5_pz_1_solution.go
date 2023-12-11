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

func get_mapping_value(input_value int, string_map string) int {
	data := strings.Split(string_map, ":\n")[1]
	conversions := strings.Split(data, "\n")
	for _, rule := range conversions {
		ranges := strings.Fields(rule)
		var _ranges = make([]int, 3)
		for i := range ranges {
			val, err := strconv.Atoi(ranges[i])
			_ranges[i] = val
			check(err)
		}
		diff := input_value - _ranges[1]
		if diff >= 0 && diff < _ranges[2] {
			return _ranges[0] + diff
		}
	}
	return input_value
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	input := string(dat)
	tables_list := strings.Split(input, "\n\n")

	lowest_location := 1e20

	seeds := strings.Split(tables_list[0], ": ")[1]
	seeds_list := strings.Fields(seeds)

	for _, seed := range seeds_list {
		_seed, err := strconv.Atoi(seed)
		check(err)
		for _, table := range tables_list[1:] {
			_seed = get_mapping_value(_seed, table)
		}
		if float64(_seed) < lowest_location {
			lowest_location = float64(_seed)
		}
	}
	fmt.Println("Result is:")
	fmt.Println(int(lowest_location))
}
