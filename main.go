package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// On recupere l'entrée  de l'utilisateur
	input := os.Getenv("INPUT_TEXT")
	if input == "" {
		fmt.Println("Aucun texte fourni")
		return
	}

	// puis le texte en minuscules
	lowercase := strings.ToLower(input)

	// et on  affiche le résultat
	fmt.Println("Texte en minuscules:", lowercase)
}
