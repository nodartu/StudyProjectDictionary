package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var FileName string
	fmt.Scan(&FileName)
	taboo := make(map[string]struct{})

	file, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ReadFile(file, taboo)
	CheckWord(taboo)

}

func ReadFile(file *os.File, taboo map[string]struct{}) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		taboo[scanner.Text()] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func CheckWord(taboo map[string]struct{}) {
	var CheckIt string
	for {
		fmt.Scan(&CheckIt)
		LowCheck := strings.ToLower(CheckIt)
		if LowCheck == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		} else if _, ok := taboo[LowCheck]; ok {
			for i := 0; i < len(LowCheck); i++ {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		} else {
			fmt.Print(CheckIt, " ")
		}
	}
}
