package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("File reading error:", err)
		return
	}
	defer file.Close()

	data := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Integer conversion error:", err)
		}
		data = append(data, integer)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanning error:", err)
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			sum := data[i] + data[j]

			if sum == 2020 {
				fmt.Println("Found one!")
				fmt.Println(data[i], data[j])
				fmt.Println(data[i] * data[j])
				return
			}
		}
	}
}
