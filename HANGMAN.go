package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Cyan    = "\033[36m"
	Magenta = "\033[35m"
)

var hangmanStages = []string{
	` 
   ----- 
   |   | 
       | 
       | 
       | 
       | 
=========`,

	// Ajoutez les autres étapes ici...
}

func main() {
	// Menu de sélection de niveau
	var level int
	fmt.Println(Green + "Bienvenue au jeu de pendu !" + Reset)
	fmt.Println("Choisissez un niveau de difficulté :")
	fmt.Println("1. Facile (10 essais)")
	fmt.Println("2. Moyen (7 essais)")
	fmt.Println("3. Difficile (5 essais)")
	fmt.Print("Votre choix : ")
	fmt.Scan(&level)

	// Définir le nombre d'essais selon le niveau
	var attempts int
	switch level {
	case 1:
		attempts = 10
	case 2:
		attempts = 7
	case 3:
		attempts = 5
	default:
		fmt.Println(Red + "Niveau invalide, vous serez placé au niveau facile par défaut." + Reset)
		attempts = 10
	}

	words, err := loadWords("mot.txt")
	if err != nil {
		fmt.Println(Red+"Erreur lors du chargement des mots :"+Reset, err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	var guessedLetters []rune

	fmt.Printf(Cyan+"Vous avez %d tentatives.\n"+Reset, attempts)

	for attempts > 0 {
		fmt.Println()
		displayHangman(attempts)
		displayWord(word, guessedLetters)

		var input string
		fmt.Print(Yellow + "Entrez une lettre ou un mot entier : " + Reset)
		fmt.Scanln(&input)

		if len(input) == 1 {
			if !containsRune(guessedLetters, rune(input[0])) {
				guessedLetters = append(guessedLetters, rune(input[0]))
				if !strings.Contains(word, input) {
					attempts--
					fmt.Println(Red + "Mauvaise lettre !" + Reset)
				} else {
					fmt.Println(Green + "Bonne lettre !" + Reset)
				}
			} else {
				fmt.Println(Magenta + "Vous avez déjà deviné cette lettre." + Reset)
			}
		} else if input == word {
			fmt.Println(Green+"Bravo ! Vous avez deviné le mot :"+Reset, word)
			return
		} else {
			attempts -= 2
			fmt.Println(Red + "Mauvais mot !" + Reset)
		}

		if isWordGuessed(word, guessedLetters) {
			fmt.Println(Green+"Félicitations ! Vous avez deviné le mot :"+Reset, word)
			return
		}

		fmt.Printf(Cyan+"Tentatives restantes : %d\n"+Reset, attempts)
	}

	fmt.Println(Red+"Dommage ! Vous avez perdu. Le mot était :"+Reset, word)
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

func displayHangman(attempts int) {
	index := len(hangmanStages) - attempts - 1
	if index >= 0 && index < len(hangmanStages) {
		fmt.Println(hangmanStages[index])
	}
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
