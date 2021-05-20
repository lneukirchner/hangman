package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func clear() {
	cmdName := "clear"
	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	triesPtr := flag.Int("tries", 6, "The amount of tries that the answerer has.")

	fmt.Print("Type your word in lowercase: ")
	var word string
	fmt.Scan(&word)
	clear()

	var tries int = *triesPtr
	var answer []string
	var triedLetters []string
	var j int
	var i int
	var k int
	for j = 0; j < len(word); j++ {
		answer = append(answer, "_")
	}
	for i = 0; i < tries; i++ {
		fmt.Print("You've tried ", triedLetters, ".\n")
		fmt.Println("You have", tries-i, "left.")
		fmt.Println(answer)
		var letter string
		fmt.Scanln(&letter)
		triedLetters = append(triedLetters, letter)
		asciiValue := letter[0]
		for k = 0; k < len(word); k++ {
			if asciiValue == word[k] {
				answer[k] = letter
			}
		}
		var total int
		for j = 0; j < len(word); j++ {
			if answer[j][0] == word[j] {
				total += 1
			} else {
				break
			}
		}
		if total == len(word) {
			break
		}
		clear()
	}
	clear()
	var total int
	for j = 0; j < len(word); j++ {
		if answer[j][0] == word[j] {
			total += 1
		} else {
			break
		}
	}
	if total == len(word) {
		fmt.Println("You Won! The word was", word, ".")
	} else {
		fmt.Println("You lost, the word was", word, ".")
	}
}
