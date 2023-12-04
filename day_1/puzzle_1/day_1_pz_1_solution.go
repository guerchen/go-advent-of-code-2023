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

func reverse(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main(){
	dat, err := os.ReadFile("../input.txt")
    check(err)
	input := string(dat)
	input_list := strings.Split(input, "\n")
	sum_calibrations := 0
	for _, element := range input_list {
        code := strings.Split(element, "")
		first_digit := "0"
		last_digit := "0"
		for _, item := range code {
			if strings.ContainsAny(item, "0123456789") {
		 		first_digit = item
		 		break
		 	}
		}
		for _, item := range reverse(code) {
		 	if strings.ContainsAny(item, "0123456789") {
		 		last_digit = item
		 		break
		 	}
		}
		calibration_value := first_digit + last_digit
		num, err := strconv.Atoi(calibration_value)
		check(err)
		sum_calibrations += num
    }
	fmt.Println("Result is:")
    fmt.Println(sum_calibrations)
}