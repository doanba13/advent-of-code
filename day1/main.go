package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	max := 0
	elfCalo := 0

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if max < elfCalo {
				max = elfCalo
			}
			elfCalo = 0
		} else {
			calo, _ := strconv.Atoi(line)
			elfCalo += calo
		}

	}

	fmt.Println(max)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
