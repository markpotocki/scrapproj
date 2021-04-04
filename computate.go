package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x int
	// given an input of accuracy A
	// find e^x where
	// = 1 + x/1! + x^2/2! + ... (A)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter accuracy: ")
	accuracy := readInt(reader)
	fmt.Print("Enter exponent: ")
	x = readInt(reader)
	value := multiThreadCompute(x, accuracy)
	fmt.Println("value is", value)
}

func iterativeCompute(x int, accuracy int) float64 {
	sum := float64(1)
	for i := 1; i <= accuracy; i++ {
		sum += computeValue(x, accuracy)
	}
	return sum
}

func multiThreadCompute(x int, accuracy int) float64 {
	sumChan := make(chan float64, accuracy)
	for i := 1; i <= accuracy; i++ {
		go func(position int) {
			sumChan <- computeValue(x, position)
		}(i)
	}

	value := float64(1)
	for i := 1; i <= accuracy; i++ {
		value += <-sumChan
	}
	return value
}

func computeValue(x int, position int) float64 {
	return math.Pow(float64(x), float64(position)) / float64(factorial(position))
}

func factorial(limit int) int {
	val := 1
	for i := limit; i > 0; i-- {
		val = val * i
	}
	return val
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readInt(reader *bufio.Reader) int {
	in, err := reader.ReadString('\n')
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	checkError(err)
	n, err := strconv.ParseInt(fmt.Sprint(in), 10, 0)
	checkError(err)
	return int(n)
}
