package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("useage : <calc> <op1> <op2>")
		os.Exit(1)
	}

	q1 := os.Args[1]
	a2 := os.Args[2]
	a3 := os.Args[3]

	q2, err := strconv.Atoi(a2)
	if err != nil {
		fmt.Println("enter valid integer")
		os.Exit(1)
	}

	q3, err := strconv.Atoi(a3)
	if err != nil {
		fmt.Println("enter valid integer")
		os.Exit(1)
	}

	switch strings.ToLower(q1) {
	case "add":
		fmt.Print(q2 + q3)
	case "sub":
		fmt.Print(q2 - q3)
	case "mul":
		fmt.Print(q2 * q3)
	case "div":
		if q3 == 0 {
			fmt.Print("enter appropriate denominator")
		} else {
			fmt.Print(q2 / q3)
		}
	}
}
