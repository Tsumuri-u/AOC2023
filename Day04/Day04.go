package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func part_one(data []string) {
	sum := 0
	for _, c := range data {
		winning := strings.Fields(strings.Split((strings.Split(c, ":")[1]), "|")[0])
		mine := strings.Fields(strings.Split((strings.Split(c, ":")[1]), "|")[1])
		if i := len(intersection(winning, mine)); i >= 1 {
			sum += int(math.Pow(float64(2), float64(i-1)))
		}
	}
	fmt.Println(sum)
}

func part_two(data []string) {
	copies := make([]int, len(data))
	for i := range data {
		copies[i] = 1
	}
	for i, c := range data {
		winning := strings.Fields(strings.Split((strings.Split(c, ":")[1]), "|")[0])
		mine := strings.Fields(strings.Split((strings.Split(c, ":")[1]), "|")[1])
		if matches := len(intersection(winning, mine)); matches != 0 {
			for j := 1; j <= matches; j++ {
				copies[i+j] += copies[i]
			}
		}
	}
	sum := 0
	for _, c := range copies {
		sum += c
	}
	fmt.Println(sum)
}

func main() {
	content, _ := os.ReadFile("input04.txt")
	data := strings.Split(string(content), "\n")
	part_one(data)
	part_two(data)
}
