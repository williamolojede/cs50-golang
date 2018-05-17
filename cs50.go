package cs50

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v", msg)
	x, _ := reader.ReadString('\n')
	x = x[:len(x)-1]
	return x
}

func retry() string {
	return getInput("Retry: ")
}

func GetFloat(msg string) (f float64, err error) {
	x := getInput(msg)
	f, err = strconv.ParseFloat(x, 32)
	return
}

func GetInt(msg string) int {
	x := getInput(msg)
	num, err := strconv.Atoi(x)

	for err != nil {
		x = retry()
		num, err = strconv.Atoi(x)
	}

	return num
}

func GetString(msg string) (str string) {
	str = getInput(msg)
	for str == "" {
		str = retry()
	}
	return
}

func GetChar(msg string) (str string) {
	str = getInput(msg)
	for len(str) > 1 || str == "" {
		str = retry()
	}
	return
}
