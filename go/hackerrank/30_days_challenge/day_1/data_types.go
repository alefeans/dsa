package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	iMy, _ := strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	dMy, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	sMy := scanner.Text()

	// Output
	fmt.Println(iMy + i)
	fmt.Printf("%.1f\n", dMy+d)
	fmt.Println(s + sMy)

}
