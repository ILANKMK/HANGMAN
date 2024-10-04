package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	words, err := loadWords("mot.txt")
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots :", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	attempts := 10
	var guessedLetters []rune

	fmt.Println("Bienvenue au jeu de pendu !")
	fmt.Printf("Vous avez %d tentatives.\n", attempts)

	for attempts > 0 {
		displayWord(word, guessedLetters)

		var input string
		fmt.Print("Entrez une lettre ou un mot entier : ")
		fmt.Scanln(&input)

		if len(input) == 1 {
			if !containsRune(guessedLetters, rune(input[0])) {
				guessedLetters = append(guessedLetters, rune(input[0]))
				if !strings.Contains(word, input) {
					attempts--
					fmt.Println("Mauvaise lettre !")
				}
			} else {
				fmt.Println("Vous avez déjà deviné cette lettre.")
			}
		} else {
			if input == word {
				fmt.Println("Bravo ! Vous avez deviné le mot :", word)
				return
			} else {
				attempts -= 2
				fmt.Println("Mauvais mot !")
			}
		}

		if isWordGuessed(word, guessedLetters) {
			fmt.Println("Félicitations ! Vous avez deviné le mot :", word)
			return
		}

		fmt.Printf("Tentatives restantes : %d\n", attempts)
	}

	fmt.Println("Dommage ! Vous avez perdu. Le mot était :", word)
}

func loadWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func displayWord(word string, guessedLetters []rune) {
	for _, letter := range word {
		if containsRune(guessedLetters, letter) {
			fmt.Print(string(letter) + " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func isWordGuessed(word string, guessedLetters []rune) bool {
	for _, letter := range word {
		if !containsRune(guessedLetters, letter) {
			return false
		}
	}
	return true
}

func containsRune(runes []rune, r rune) bool {
	for _, v := range runes {
		if v == r {
			return true
		}
	}
	return false
}
