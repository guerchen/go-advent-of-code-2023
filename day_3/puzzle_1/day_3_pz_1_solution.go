//I couldn't solve day 3. My solution works on test dataset, but didn't on real input. Can't figure why. Moving to day 4.

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

type point struct {
	x int
	y int
}

func get_neighbors(p point, schema [][]string) [][]int {
	var neighbors [][]int
	diffs := []int{-1, 0, 1}

	for _, i := range diffs {
		for _, j := range diffs {
			if i == 0 && j == 0 {
				continue
			}
			if p.x+i < 0 || p.y+j < 0 {
				continue
			}
			if p.x+i >= len(schema) || p.y+j >= len(schema[0]) {
				continue
			}
			neighbor_coords := []int{p.x + i, p.y + j}
			neighbors = append(neighbors, neighbor_coords)
		}
	}

	return neighbors
}

func extract_number(p point, schema [][]string) (int, int) {
	min_delta, max_delta := 0, 0

	for (p.y+min_delta-1 >= 0) && strings.ContainsAny(schema[p.x][p.y+min_delta-1], "0123456789") {
		min_delta -= 1
	}

	for (p.y+max_delta+1 < len(schema[0])) && strings.ContainsAny(schema[p.x][p.y+max_delta+1], "0123456789") {
		max_delta += 1
	}
	str_number_array := strings.Join(schema[p.x][p.y+min_delta:p.y+max_delta+1], "")
	part_num, err := strconv.Atoi(str_number_array)
	check(err)

	return part_num, max_delta
}

func is_adjacent(neighbors [][]int, schema [][]string) bool {
	adjacent := false
	for _, neighbor := range neighbors {
		if !(strings.ContainsAny(schema[neighbor[0]][neighbor[1]], "0123456789")) &&
			(schema[neighbor[0]][neighbor[1]] != ".") {
			adjacent = true
		}
	}
	return adjacent
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	input := string(dat)

	schema := strings.Split(input, "\n")

	schema_array := make([][]string, len(schema))

	for i, schema_line := range schema {
		schema_array[i] = strings.Split(schema_line, "")
	}

	part_number_sum := 0

	for i, row := range schema_array {
		skip := 0
		for j := range row {
			if skip == 0 {
				if strings.ContainsAny(schema_array[i][j], "0123456789") {
					current_coords := point{i, j}
					neighbors := get_neighbors(current_coords, schema_array)
					if is_adjacent(neighbors, schema_array) {
						num, _skip := extract_number(current_coords, schema_array)
						//fmt.Println(num)
						skip = _skip
						part_number_sum += num
						//fmt.Println(part_number_sum)
					}
				}
			} else {
				skip -= 1
			}
		}
	}
	fmt.Println("Result is:")
	fmt.Println(part_number_sum)
}
