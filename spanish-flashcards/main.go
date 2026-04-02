package main

import (
	"fmt"
	"strings"

	//"strings"
	//"io/ioutil"
	"os"
)

func readfile() string {
	data, err := os.ReadFile("words.txt")

	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	data := readfile()

	arrData := strings.Split(data, ",")
	flashCard := make(map[string]string)

	for _, val := range arrData {
		val2 := strings.Split(val, ":")
		flashCard[val2[0]] = val2[1]

	}

	fmt.Print(flashCard)

}
