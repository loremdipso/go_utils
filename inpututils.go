package go_utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInt() int {
	var tmp int
	fmt.Scan(&tmp)
	return tmp
}

func GetFloat() float64 {
	var tmp float64
	fmt.Scan(&tmp)
	return tmp
}

func GetString() string {
	var tmp string
	fmt.Scan(&tmp)
	return tmp
}

func GetLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return line
}

func ReadChar(prompt string, interfaces ...interface{}) (string, error) {
	fmt.Printf(prompt, interfaces...)
	reader := bufio.NewReader(os.Stdin)
	r, _, err := reader.ReadRune()
	return string(r), err
}

// TODO: this might... just be wrong
func Readline(prompt string, interfaces ...interface{}) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt, interfaces...)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.Trim(input, "\n\r")
	return input, nil
}
