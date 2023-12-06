package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_numbers(s string) int {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, n := range numbers {
		if strings.Contains(s, n) {
			return i + 1
		}
	}
	return 0
}

func print_sum(nums []int) {
	sum := 0
	for _, c := range nums {
		sum += c
	}
	fmt.Println(sum)
}

func part_one(arr []string) {
	var values []int

	for _, s := range arr {
		var temp []int
		for _, c := range s {
			if i, err := strconv.Atoi(string(c)); err == nil {
				temp = append(temp, i)
			}
		}
		values = append(values, (temp[0]*10 + temp[len(temp)-1]))
	}
	print_sum(values)
}

func part_two(arr []string) {
	var values []int

	for _, s := range arr {
		var digits []int
		strings := ""

		for _, c := range s {
			if i, err := strconv.Atoi(string(c)); err == nil {
				digits = append(digits, i)
			} else {
				strings = strings + string(c)
				if num := check_numbers(strings); num != 0 {
					digits = append(digits, num)
					strings = string(c)
				}
			}
		}
		values = append(values, (digits[0]*10 + digits[len(digits)-1]))
	}
	print_sum(values)
}

func main() {
	content, _ := os.ReadFile("input01.txt")
	data := strings.Split(string(content), "\n")
	part_one(data)
	part_two(data)
}
