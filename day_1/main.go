package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to read input: " + err.Error())
	}

	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		local_sum := ""

		for _, c := range str {
			if isNumber(string(c)) {
				local_sum += string(c)
			}
		}

		if len(local_sum) > 2 {
			local_sum = string(local_sum[0]) + string(local_sum[len(local_sum)-1])
		}

		if len(local_sum) == 1 {
			local_sum = local_sum + local_sum
		}

		i, err := strconv.Atoi(local_sum)

		if err != nil {
			log.Fatal(err)
		}

		sum += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum:", sum)
}

func isNumber(str string) bool {
	arr := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
	}

	return false
}
