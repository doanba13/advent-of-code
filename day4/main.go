package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func priorityOf(s byte) int {
	test := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(test, string(s)) + 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	containRangeCount := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		e1 := strings.Split(line[0], "-")
		e2 := strings.Split(line[1], "-")

		e1Start, _ := strconv.Atoi(e1[0])
		e1Stop, _ := strconv.Atoi(e1[1])

		e2Start, _ := strconv.Atoi(e2[0])
		e2Stop, _ := strconv.Atoi(e2[1])

		if (e1Start <= e2Start && e1Stop >= e2Stop) || (e1Start >= e2Start && e1Stop <= e2Stop) {
			containRangeCount++
		}
	}

	fmt.Println(containRangeCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
