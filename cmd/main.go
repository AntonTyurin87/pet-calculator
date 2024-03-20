package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pet_calculator/pkg/add"
	"strings"
)

func main() {
	fmt.Println("Выберете операцию из списка:")
	fmt.Println("SUM или 1")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)

	switch {
	case input == "SUM" || input == "1":
		add.Add()
	}

}
