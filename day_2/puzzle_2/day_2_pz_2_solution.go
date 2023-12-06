package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func power_cubes(rounds []string) int {
	min_red, min_green, min_blue := 0, 0, 0

	for _, round := range rounds {
		draw := strings.Split(round, ", ")
		for _, cube := range draw {
			cube_info := strings.Split(cube, " ")
			cube_amount, err := strconv.Atoi(cube_info[0])
			check(err)
			cube_color := cube_info[1]

			if cube_color == "red" && cube_amount > min_red {
				min_red = cube_amount
			} 
			if cube_color == "green" && cube_amount > min_green {
				min_green = cube_amount
			}
			if cube_color == "blue" && cube_amount > min_blue {
				min_blue = cube_amount
			}
		}
	}

	return min_red * min_green * min_blue
}

func main(){
	dat, err := os.ReadFile("../input.txt")
    check(err)
	input := string(dat)

	games_list := strings.Split(input, "\n")
	games_list = games_list[:len(games_list)-1]

	sum_powers := 0
	for _, element := range games_list {

		game_info := strings.Split(element, ": ")

		rounds_list := strings.Split(game_info[1], "; ")

		sum_powers += power_cubes(rounds_list)

    }
	fmt.Println("Result is:")
    fmt.Println(sum_powers)
}