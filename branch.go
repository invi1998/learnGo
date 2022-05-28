package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func grade(source int) string {
	g := ""
	switch {
	case source < 0 || source > 100:
		panic(fmt.Sprintf("Wrong source: %d", source))
	case source < 60:
		g = "F"
	case source <= 80:
		g = "B"
	case source <= 90:
		g = "A"
	case source <= 100:
		g = "A"
	}

	return g
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forover() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	const filename = "abc.txt"
	if constens, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constens)
	}

	fmt.Println(grade(0))
	fmt.Println(grade(60))
	fmt.Println(grade(70))
	fmt.Println(grade(80))
	fmt.Println(grade(90))
	//fmt.Println(grade(100))
	//fmt.Println(grade(120))

	fmt.Println(
		convertToBin(5),        // 101
		convertToBin(13),       // 1101
		convertToBin(72387885), // 100010100001000110100101101
		convertToBin(0),
	)

	printFile("abc.txt")

	forover()
}
