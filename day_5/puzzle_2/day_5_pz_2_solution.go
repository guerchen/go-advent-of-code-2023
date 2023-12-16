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

func build_seed_range(seeds_range_list []string) []int {
	seed_list := []int{}

	var _seeds_range_list = make([]int, len(seeds_range_list))
	for i := range seeds_range_list {
		val, err := strconv.Atoi(seeds_range_list[i])
		_seeds_range_list[i] = val
		check(err)
	}

	for i := range _seeds_range_list {
		if i%2 != 0 {
			continue
		} else {
			for j := _seeds_range_list[i]; j < _seeds_range_list[i+1]; j++ {
				seed_list = append(seed_list, j)
			}
		}
	}

	return seed_list
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
	seeds_range_list := strings.Fields(seeds)

	seeds_list := build_seed_range(seeds_range_list)

	for _, seed := range seeds_list {
		for _, table := range tables_list[1:] {
			seed = get_mapping_value(seed, table)
		}
		if float64(seed) < lowest_location {
			lowest_location = float64(seed)
		}
	}
	fmt.Println("Result is:")
	fmt.Println(int(lowest_location))
}
