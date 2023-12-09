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

type numbersSet map[string]bool

func instantiate(numberArray []string) numbersSet {
	var newSet = make(numbersSet)
	for _, num := range numberArray {
		if num != " " {
			newSet.add(num)
		}
	}
	return newSet
}

func (s numbersSet) add(number string) {
	s[number] = true
}

func (s numbersSet) has(number string) bool {
	_, ok := s[number]
	return ok
}

func (s numbersSet) intersection(otherSet map[string]bool) map[string]bool {
	var newSet = make(numbersSet)
	for key := range otherSet {
		if s.has(key) {
			newSet.add(key)
		}
	}
	return newSet
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	input := string(dat)
	games_list := strings.Split(input, "\n")

	sum_cards := 0
	num_cards := make(map[int]int)
	for i := 1; i <= len(games_list); i++ {
		num_cards[i] = 1
	}

	for _, element := range games_list {

		game_info := strings.Split(element, ": ")
		game_num := strings.Fields(game_info[0])[1]
		_game_num, err := strconv.Atoi(game_num)
		check(err)

		numbers := strings.Split(game_info[1], " | ")

		winning_numbers := strings.Fields(numbers[0])
		number_owned := strings.Fields(numbers[1])

		winning_set := instantiate(winning_numbers)
		owned_set := instantiate(number_owned)

		inter := winning_set.intersection(owned_set)
		points := len(inter)

		for i := _game_num + 1; i <= _game_num+points; i++ {
			num_cards[i] += 1 * num_cards[_game_num]
		}

		sum_cards += num_cards[_game_num]

	}
	fmt.Println("Result is:")
	fmt.Println(sum_cards)
}
