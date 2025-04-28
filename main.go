package main

import (
	"github.com/FrancoMusolino/film-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}
