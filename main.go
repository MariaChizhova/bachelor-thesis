package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("benchmarks/2.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`\d+\.?\d+? ns/op`)
	results := []float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindString(line)

		if match != "" {
			value, err := strconv.ParseFloat(match[:len(match)-6], 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			results = append(results, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(strings.Join(strings.Fields(fmt.Sprint(results)), ","))
}
