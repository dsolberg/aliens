package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("maps/citymap.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text())  // token in unicode-char
			fmt.Println(scanner.Bytes()) // token in bytes
	}

}
