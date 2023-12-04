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

func assign_value(first_digit string, last_digit string, found_digit string) (string, string) {
	if first_digit == "0" {
		return found_digit, found_digit
	} else {
		return first_digit, found_digit
	}
}

func find_digits(code []string) (string,string) {
	first_digit, last_digit := "0", "0"
	for idx, item := range code {
		if strings.ContainsAny(item, "0123456789") {
			first_digit, last_digit = assign_value(first_digit, last_digit, item)
		} 
		if len(code) - idx >= 3 {
			if strings.Join(code[idx:idx+3], "") == "one" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "1")		
			} else if strings.Join(code[idx:idx+3], "") == "two" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "2")			
			} else if strings.Join(code[idx:idx+3], "") == "six" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "6")			
			}
		}
		if len(code) - idx >= 4 {
			if strings.Join(code[idx:idx+4], "") == "four" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "4")
			} else if strings.Join(code[idx:idx+4], "") == "five" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "5")			
			} else if strings.Join(code[idx:idx+4], "") == "nine" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "9")
			} else if strings.Join(code[idx:idx+4], "") == "zero" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "0")
			}
		}
		if len(code) - idx >= 5 {
			if strings.Join(code[idx:idx+5], "") == "three" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "3")			
			} else if strings.Join(code[idx:idx+5], "") == "seven" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "7")			
			} else if strings.Join(code[idx:idx+5], "") == "eight" {
				first_digit, last_digit = assign_value(first_digit, last_digit, "8")			
			}
		}
		  
	}
	return first_digit, last_digit
}



func main(){
	dat, err := os.ReadFile("../input.txt")
    check(err)
	input := string(dat)
	input_list := strings.Split(input, "\n")
	sum_calibrations := 0
	for _, element := range input_list {
        code := strings.Split(element, "")
		first_digit, last_digit := find_digits(code)
		
		calibration_value := first_digit + last_digit
		num, err := strconv.Atoi(calibration_value)
		check(err)
		sum_calibrations += num
    }
	fmt.Println("Result is:")
    fmt.Println(sum_calibrations)
}