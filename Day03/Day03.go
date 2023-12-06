package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_val(data byte, tempnum *int, left bool, order int) bool {
	if n, err := strconv.Atoi(string(data)); err == nil {
		if left {
			*tempnum = n*order + *tempnum
		} else {
			*tempnum = *tempnum*order + n
		}
		return true
	}
	return false
}

func part_one(data []string) {
	sum := 0
	for i, line := range data {
		tempnum := 0
		part := false

		for j, char := range line {
			if int, err := strconv.Atoi(string(char)); err == nil {
				tempnum = (tempnum * 10) + int
				left, right, top, bottom := max(0, j-1), min(len(data[0])-1, j+1), max(0, i-1), min(len(data)-1, i+1)
				border := []byte{data[top][left], data[top][j], data[top][right], data[i][left],
					data[i][right], data[bottom][left], data[bottom][j], data[bottom][right]}

				if strings.ContainsAny(string(border), "*/=%@&#$-+") {
					part = true
				}
				if j == len(data[0])-1 && part {
					sum += tempnum
				}
			} else {
				if part {
					sum += tempnum
					part = false
				}
				tempnum = 0
			}
		}
	}
	fmt.Println(sum)
}

func part_two(data []string) {
	sum := 0
	for i, line := range data {
		for j, char := range line {
			if string(char) == "*" {
				tempnum := 0
				actualnum := 0
				topmiddle := false
				bottommiddle := false
				// check top middle
				if check_val(data[i-1][j], &tempnum, false, 0) {
					if check_val(data[i-1][j-1], &tempnum, true, 10) {
						check_val(data[i-1][j-2], &tempnum, true, 100)
					}
					if check_val(data[i-1][j+1], &tempnum, false, 10) {
						check_val(data[i-1][j+2], &tempnum, false, 10)
					}
					topmiddle = true
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				// check top left
				if !topmiddle {
					if check_val(data[i-1][j-1], &tempnum, false, 0) {
						if check_val(data[i-1][j-2], &tempnum, true, 10) {
							check_val(data[i-1][j-3], &tempnum, true, 100)
						}
					}
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				// check top right
				if !topmiddle {
					if check_val(data[i-1][j+1], &tempnum, false, 0) {
						if check_val(data[i-1][j+2], &tempnum, false, 10) {
							check_val(data[i-1][j+3], &tempnum, false, 10)
						}
					}
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				// check left
				if check_val(data[i][j-1], &tempnum, false, 0) {
					if check_val(data[i][j-2], &tempnum, true, 10) {
						check_val(data[i][j-3], &tempnum, true, 100)
					}
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				// check right
				if check_val(data[i][j+1], &tempnum, false, 0) {
					if check_val(data[i][j+2], &tempnum, false, 10) {
						check_val(data[i][j+3], &tempnum, false, 10)
					}
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				// check bottom middle
				if check_val(data[i+1][j], &tempnum, false, 0) {
					if check_val(data[i+1][j-1], &tempnum, true, 10) {
						check_val(data[i+1][j-2], &tempnum, true, 100)
					}
					if check_val(data[i+1][j+1], &tempnum, false, 10) {
						check_val(data[i+1][j+2], &tempnum, false, 10)
					}
					bottommiddle = true
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				// check bottom left
				if !bottommiddle {
					if check_val(data[i+1][j-1], &tempnum, false, 0) {
						if check_val(data[i+1][j-2], &tempnum, true, 10) {
							check_val(data[i+1][j-3], &tempnum, true, 100)
						}
					}
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				// check bottom right
				if !bottommiddle {
					if check_val(data[i+1][j+1], &tempnum, false, 0) {
						if check_val(data[i+1][j+2], &tempnum, false, 10) {
							check_val(data[i+1][j+3], &tempnum, false, 10)
						}
					}
					if actualnum == 0 {
						actualnum = tempnum
						tempnum = 0
					}
				}
				sum += tempnum * actualnum
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	content, _ := os.ReadFile("input03.txt")
	data := strings.Split(string(content), "\n")
	part_one(data)
	part_two(data)
}
