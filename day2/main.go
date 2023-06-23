package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getScore(first, last string, d map[string](int)) int {
	if last == "X" {
		score := d[first] - 1
		if score == 0 {
			return 3
		}
		return score
	}

	if last == "Y" {
		return d[first] + 3
	}

	fmt.Printf("%s - %s \n", last, first)

	score := d[first] + 1
	if score == 4 {
		return 6 + 1
	}
	return score + 6
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dics := make(map[string](int))

	dics["A"] = 1 // Rock
	dics["B"] = 2 // Paper
	dics["C"] = 3 // Scissor

	dics["X"] = 1
	dics["Y"] = 4
	dics["Z"] = 7

	totalScore := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, " ")
		a := getScore(l[0], l[1], dics)

		fmt.Println(a)

		totalScore += a
	}

	fmt.Println(totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
