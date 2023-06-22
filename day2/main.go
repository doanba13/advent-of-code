package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getScore(first, last string, d map[string](int)) int {
	c := d[last] - d[first]

	fmt.Printf("%s - %s = %d \n", last, first, c)

	if c == 1 {
		return d[last] + 6
	}

	if c == 0 {
		return d[last] + 3
	}

	if c == -2 {
		return d[last] + 6
	}

	if c == 2 {
		return d[last]
	}

	return d[last]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dics := make(map[string](int))

	dics["A"] = 1
	dics["B"] = 2
	dics["C"] = 3

	dics["X"] = 1
	dics["Y"] = 2
	dics["Z"] = 3

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
