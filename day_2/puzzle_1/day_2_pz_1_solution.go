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

func check_possibility(rounds []string) bool {

	for _, round := range rounds {
		draw := strings.Split(round, ", ")
		for _, cube := range draw {
			cube_info := strings.Split(cube, " ")
			cube_amount, err := strconv.Atoi(cube_info[0])
			check(err)
			cube_color := cube_info[1]

			if cube_color == "red" && cube_amount > 12 {
				return false
			} 
			if cube_color == "green" && cube_amount > 13 {
				return false
			}
			if cube_color == "blue" && cube_amount > 14 {
				return false
			}
		}
	}

	return true
}

func main(){
	dat, err := os.ReadFile("../input.txt")
    check(err)
	input := string(dat)
	games_list := strings.Split(input, "\n")
	games_list = games_list[:len(games_list)-1]
	sum_games := 0
	for _, element := range games_list {

		game_info := strings.Split(element, ": ")
		game_id, err := strconv.Atoi(game_info[0][5:])
		check(err)

		rounds_list := strings.Split(game_info[1], "; ")

		if (check_possibility(rounds_list)) {
			sum_games += game_id
		}

    }
	fmt.Println("Result is:")
    fmt.Println(sum_games)
}