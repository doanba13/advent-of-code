package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	itemPrioritySum := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstCompartment := line[0 : len(line)/2]
		secondCompartment := line[len(line)/2:]

		for i := 0; i < len(firstCompartment); i++ {
			l := string(firstCompartment[i])
			if strings.Contains(firstCompartment, l) && strings.Contains(secondCompartment, l) {
				fmt.Println(string(firstCompartment[i]))
				itemPrioritySum += priorityOf(firstCompartment[i])

				break
			}
		}
	}

	fmt.Println(itemPrioritySum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
