package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var max []int
	elfCalo := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			max = append(max, elfCalo)
			elfCalo = 0
		} else {
			calo, _ := strconv.Atoi(line)
			elfCalo += calo
		}
	}

	sort.Slice(max, func(i, j int) bool {
		return max[i] > max[j]
	})

	sum := 0

	for i := 0; i < 3; i++ {
		sum += max[i]
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
