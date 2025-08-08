package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makePhoneBook(scanner *bufio.Scanner, n int) map[string]string {
	phoneBook := make(map[string]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		entry := strings.Split(scanner.Text(), " ")
		phoneBook[entry[0]] = entry[1]
	}
	return phoneBook
}

func queries(scanner *bufio.Scanner, phoneBook map[string]string) {
	for scanner.Scan() {
		if val, ok := phoneBook[scanner.Text()]; ok {
			fmt.Printf("%s=%s\n", scanner.Text(), val)
		} else {
			fmt.Println("Not found")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	phoneBook := makePhoneBook(scanner, n)
	queries(scanner, phoneBook)
}
