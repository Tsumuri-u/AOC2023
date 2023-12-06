package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part_one(arr []string) {
	red, green, blue := 12, 13, 14
	possible := 5050
	for i, c := range arr {
		pulls := strings.FieldsFunc(c, func(r rune) bool { return r == ';' || r == ':' || r == ',' })[1:]
		for _, p := range pulls {
			num, _ := strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(p))
			if (strings.Contains(p, "red") && num > red) ||
				(strings.Contains(p, "green") && num > green) ||
				(strings.Contains(p, "blue") && num > blue) {
				possible -= (i + 1)
				break
			}
		}
	}
	fmt.Println(possible)
}

func part_two(arr []string) {
	power := 0
	for _, c := range arr {
		pulls := strings.FieldsFunc(c, func(r rune) bool { return r == ';' || r == ':' || r == ',' })[1:]
		red, green, blue := 0, 0, 0
		for _, p := range pulls {
			num, _ := strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(p))
			if strings.Contains(p, "red") && num > red {
				red = num
			} else if strings.Contains(p, "green") && num > green {
				green = num
			} else if strings.Contains(p, "blue") && num > blue {
				blue = num
			}
		}
		power += (red * green * blue)
	}
	fmt.Println(power)
}

func main() {
	content, _ := os.ReadFile("input02.txt")
	data := strings.Split(string(content), "\n")
	part_one(data)
	part_two(data)
}
