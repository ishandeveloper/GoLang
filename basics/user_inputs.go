package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name & Age : ")
	name, _ := reader.ReadString('\n')

	ageinput, _ := reader.ReadString('\n')
	age, _ := strconv.ParseFloat(strings.TrimSpace(ageinput), 64)
	fmt.Println(name)
	fmt.Println(age + 2)

}
