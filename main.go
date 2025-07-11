package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FrancoMusolino/film-cli/cmd"
	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {
	godotenv.Load(".env")
	loadTranslations()

	cmd.Execute()
}

func loadTranslations() {
	translationsDir := "./messages"
	files, err := os.ReadDir(translationsDir)
	if err != nil {
		log.Fatal("Cannot read Dir /messages")
	}

	for _, file := range files {
		i18n.MustLoadTranslationFile(fmt.Sprintf("%s/%s", translationsDir, file.Name()))
	}
}
