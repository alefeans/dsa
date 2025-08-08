package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func splitStrings(input string) (even, odd []string) {
	for j, x := range input {
		if j%2 == 0 {
			even = append(even, string(x))
		} else {
			odd = append(odd, string(x))
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		input := getString(scanner)
		even, odd := splitStrings(input)
		fmt.Println(strings.Join(even, ""), strings.Join(odd, ""))
	}
}
